package main

import (
	"archive/zip"
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"github.com/tensorflow/tensorflow/tensorflow/go/op"
)

func main() {

	// Un ejemplo para usar la API TensorFlow Go en el reconocimiento de imágenes
	// usando un modelo de inicio pre-entrenado (http://arxiv.org/abs/1512.00567).

	// ===========================================================================
	// Ejemplo de uso: <program> -dir=/tmp/modeldir -image=/path/to/someimage/jpeg
	// ===========================================================================

	// El modelo pre-entrenado toma la entrada en forma de 4 dimensiones
	// del tensor con forma [ BATCH_SIZE, IMAGE_HEIGHT, IMAGE_WIDTH, 3 ],
	// donde:
	// - BATCH_SIZE permite la inferencia de múltiples imágenes en una pasada a través del grafico
	// - IMAGE_HEIGHT es la altura de las imágenes en las que se entrenó el modelo
	// - IMAGE_WIDTH es el ancho de las imágenes en las que se entrenó el modelo
	// - 3 son los valores (R, G, B) de los colores de píxeles representados como flotantes.
	//
	// Y produce como salida un vector con forma [ NUM_LABELS ].
	// output[i] es la probabilidad de que se reconozca que la imagen
	// de entrada tiene la etiqueta i-ésima.
	//
	// Un archivo separado contiene una lista de etiquetas de cadena
	// correspondientes a los índices enteros de la salida.
	//
	// Este ejemplo:
	// - Carga la representación serializada del modelo pre-entrenado en un gráfico
	// - Crea una sesión para ejecutar operaciones en el gráfico
	// - Convierte un archivo de imagen en un Tensor para proporcionar como entrada a una sesión ejecutada
	// - Ejecuta la sesión e imprime la etiqueta con la mayor probabilidad
	//
	// Para convertir un archivo de imagen a un Tensor adecuado para ingresar
	// al modelo Inception, este ejemplo:
	// - Construye otro gráfico TensorFlow para normalizar la imagen en una
	//   forma adecuada para el modelo (por ejemplo, cambiar el tamaño de la imagen)
	// - Crea y ejecuta una sesión para obtener un tensor en esta forma normalizada.

	modeldir := flag.String("dir", "", "Directorio que contiene los archivos del modelo entrenado. Se creará el directorio y se descargará el modelo si es necesario.")
	imagefile := flag.String("image", "", "Ruta de una imagen JPEG para extraer etiquetas para")
	flag.Parse()
	if *modeldir == "" || *imagefile == "" {
		flag.Usage()
		return
	}
	// Cargar el GraphDef serializado desde un archivo.
	modelfile, labelsfile, err := modelFiles(*modeldir)
	if err != nil {
		log.Fatal(err)
	}
	model, err := ioutil.ReadFile(modelfile)
	if err != nil {
		log.Fatal(err)
	}

	// Construir gráfico en memoria a partir de formulario serializado.
	graph := tf.NewGraph()
	if err := graph.Import(model, ""); err != nil {
		log.Fatal(err)
	}

	// Crear una sesión para inferencia sobre gráfico.
	session, err := tf.NewSession(graph, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// Ejecutar inferencia en *imageFile.
	// Para varias imágenes, se puede llamar a session.Run() en un bucle (y al mismo tiempo).
	// Alternativamente, las imágenes se pueden agrupar ya que el modelo acepta lotes
	// de datos de imagen como entrada.
	tensor, err := makeTensorFromImage(*imagefile)
	if err != nil {
		log.Fatal(err)
	}
	output, err := session.Run(
		map[tf.Output]*tf.Tensor{
			graph.Operation("input").Output(0): tensor,
		},
		[]tf.Output{
			graph.Operation("output").Output(0),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}

	// output[0].Value() es un vector que contiene probabilidades de etiquetas
	// para cada imagen en el "lote". El tamaño del lote fue 1.
	// Encontrar índice de etiqueta más probable.
	probabilities := output[0].Value().([][]float32)[0]
	printBestLabel(probabilities, labelsfile)
}

func printBestLabel(probabilities []float32, labelsFile string) {
	bestIdx := 0
	for i, p := range probabilities {
		if p > probabilities[bestIdx] {
			bestIdx = i
		}
	}
	// Encontrando la mejor coincidencia. Leer cadena de labelsFile,
	// que contiene una línea por etiqueta.
	file, err := os.Open(labelsFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var labels []string
	for scanner.Scan() {
		labels = append(labels, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Printf("ERROR: failed to read %s: %v", labelsFile, err)
	}
	fmt.Printf("BEST MATCH: (%2.0f%% likely) %s\n", probabilities[bestIdx]*100.0, labels[bestIdx])
}

// Convertir la imagen en el nombre del archivo a un Tensor adecuado
// como entrada para el modelo Inception.
func makeTensorFromImage(filename string) (*tf.Tensor, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	// DecodeJpeg utiliza un tensor escalar con valor de cadena como entrada.
	tensor, err := tf.NewTensor(string(bytes))
	if err != nil {
		return nil, err
	}
	// Construir un gráfico para normalizar la imagen.
	graph, input, output, err := constructGraphToNormalizeImage()
	if err != nil {
		return nil, err
	}
	// Ejecutar ese gráfico para normalizar la imagen
	session, err := tf.NewSession(graph, nil)
	if err != nil {
		return nil, err
	}
	defer session.Close()
	normalized, err := session.Run(
		map[tf.Output]*tf.Tensor{input: tensor},
		[]tf.Output{output},
		nil)
	if err != nil {
		return nil, err
	}
	return normalized[0], nil
}

// El modelo de inicio toma como entrada la imagen descrita por un Tensor en un formato normalizado
// muy específico (un tamaño de imagen particular, forma del tensor de entrada,
// valores de píxeles normalizados, etc.).
//
// Esta función construye un gráfico de operaciones TensorFlow que toma como entrada
// una cadena codificada en JPEG y devuelve un tensor adecuado como entrada al modelo inicial.
func constructGraphToNormalizeImage() (graph *tf.Graph, input, output tf.Output, err error) {
	// Algunas constantes específicas del modelo pre-entrenado en:
	// https://storage.googleapis.com/download.tensorflow.org/models/inception5h.zip
	//
	// - El modelo fue entrenado después con imágenes escaladas a 224x224 píxeles.
	// - Los colores, representados como R,G,B en 1 byte cada uno, se convirtieron
	//   en flotantes usando (valor - Media) / Escala.
	const (
		H, W  = 224, 224
		Mean  = float32(117)
		Scale = float32(1)
	)
	// - input es un String-Tensor, donde la cadena es la imagen codificada en JPEG.
	// - El modelo inicial toma un tensor de forma 4D [BatchSize, Height, Width, Colors = 3],
	// 	 donde cada píxel se representa como un triplete de flotantes
	// - Aplicar la normalización en cada píxel y usar ExpandDims para hacer que esta única imagen
	// 	 sea un "lote" de tamaño 1 para ResizeBilinear.
	s := op.NewScope()
	input = op.Placeholder(s, tf.String)
	output = op.Div(s,
		op.Sub(s,
			op.ResizeBilinear(s,
				op.ExpandDims(s,
					op.Cast(s,
						op.DecodeJpeg(s, input, op.DecodeJpegChannels(3)), tf.Float),
					op.Const(s.SubScope("make_batch"), int32(0))),
				op.Const(s.SubScope("size"), []int32{H, W})),
			op.Const(s.SubScope("mean"), Mean)),
		op.Const(s.SubScope("scale"), Scale))
	graph, err = s.Finalize()
	return graph, input, output, err
}

func modelFiles(dir string) (modelfile, labelsfile string, err error) {
	const URL = "https://storage.googleapis.com/download.tensorflow.org/models/inception5h.zip"
	var (
		model   = filepath.Join(dir, "tensorflow_inception_graph.pb")
		labels  = filepath.Join(dir, "imagenet_comp_graph_label_strings.txt")
		zipfile = filepath.Join(dir, "inception5h.zip")
	)
	if filesExist(model, labels) == nil {
		return model, labels, nil
	}
	log.Println("Did not find model in", dir, "downloading from", URL)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", "", err
	}
	if err := download(URL, zipfile); err != nil {
		return "", "", fmt.Errorf("failed to download %v - %v", URL, err)
	}
	if err := unzip(dir, zipfile); err != nil {
		return "", "", fmt.Errorf("failed to extract contents from model archive: %v", err)
	}
	os.Remove(zipfile)
	return model, labels, filesExist(model, labels)
}

func filesExist(files ...string) error {
	for _, f := range files {
		if _, err := os.Stat(f); err != nil {
			return fmt.Errorf("unable to stat %s: %v", f, err)
		}
	}
	return nil
}

func download(URL, filename string) error {
	resp, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	return err
}

func unzip(dir, zipfile string) error {
	r, err := zip.OpenReader(zipfile)
	if err != nil {
		return err
	}
	defer r.Close()
	for _, f := range r.File {
		src, err := f.Open()
		if err != nil {
			return err
		}
		log.Println("Extracting", f.Name)
		dst, err := os.OpenFile(filepath.Join(dir, f.Name), os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		if _, err := io.Copy(dst, src); err != nil {
			return err
		}
		dst.Close()
	}
	return nil
}
