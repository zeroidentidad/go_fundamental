package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/sajari/regression"
)

// ModelInfo incluye información sobre el modelo que sale del entrenamiento.
type ModelInfo struct {
	Intercept    float64           `json:"intercept"`
	Coefficients []CoefficientInfo `json:"coefficients"`
}

// CoefficientInfo incluye información sobre un coeficiente de modelo particular.
type CoefficientInfo struct {
	Name        string  `json:"name"`
	Coefficient float64 `json:"coefficient"`
}

func main() {

	// Declarar las banderas (-flag) de directorio de entrada y salida.
	inDirPtr := flag.String("inDir", "", "The directory containing the training data")
	outDirPtr := flag.String("outDir", "", "The output directory")

	// Analizar banderas de la línea de comandos.
	flag.Parse()

	// Abrir archivo de conjunto de datos de entrenamiento.
	f, err := os.Open(filepath.Join(*inDirPtr, "diabetes.csv"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Crear nuevo lector CSV que lea del archivo abierto.
	reader := csv.NewReader(f)

	// Leer en todos los registros CSV
	reader.FieldsPerRecord = 11
	trainingData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// En este caso, se tratará de modelar medida de la enfermedad y mediante la función bmi,
	// para interceptar. Como tal, crear la estructura necesaria para entrenar un modelo
	// usando github.com/sajari/regression.
	var r regression.Regression
	r.SetObserved("diabetes progression")
	r.SetVar(0, "bmi")

	// Bucle de registros en el CSV, agregando los datos de entrenamiento al valor de regresión.
	for i, record := range trainingData {

		// Omitir encabezado
		if i == 0 {
			continue
		}

		// Analizar la medida de progresión de la diabetes, o "y".
		yVal, err := strconv.ParseFloat(record[10], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Analizar el valor de bmi.
		bmiVal, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Agregue estos puntos al valor de regresión.
		r.Train(regression.DataPoint(yVal, []float64{bmiVal}))
	}

	// Entrenar/ajustar el modelo de regresión.
	r.Run()

	// Salida de los parámetros del modelo entrenado a stdout.
	fmt.Printf("\nRegression Formula:\n%v\n\n", r.Formula)

	// Llenar en información del modelo.
	modelInfo := ModelInfo{
		Intercept: r.Coeff(0),
		Coefficients: []CoefficientInfo{
			CoefficientInfo{
				Name:        "bmi",
				Coefficient: r.Coeff(1),
			},
		},
	}

	// Marshal la información del modelo.
	outputData, err := json.MarshalIndent(modelInfo, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	// Guardar la salida ordenada en un archivo.
	if err := ioutil.WriteFile(filepath.Join(*outDirPtr, "model.json"), outputData, 0644); err != nil {
		log.Fatal(err)
	}
}
