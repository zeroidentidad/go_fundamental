package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

// neuralNet contiene toda la información
// que define una red neuronal entrenada.
type neuralNet struct {
	config  neuralNetConfig
	wHidden *mat.Dense
	bHidden *mat.Dense
	wOut    *mat.Dense
	bOut    *mat.Dense
}

// neuralNetConfig define arquitectura
// de red neuronal y parámetros de aprendizaje.
type neuralNetConfig struct {
	inputNeurons  int
	outputNeurons int
	hiddenNeurons int
	numEpochs     int
	learningRate  float64
}

func main() {

	// Abrir archivo de conjunto de datos de entrenamiento.
	f, err := os.Open("train.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Crear nuevo lector CSV que lea del archivo abierto.
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 7

	// Leer en todos los registros CSV
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// inputsData y labelsData almacenarán los valores flotantes
	// que se usarán para formar las matrices.
	inputsData := make([]float64, 4*len(rawCSVData))
	labelsData := make([]float64, 3*len(rawCSVData))

	// inputsIndex rastreará índice actual de valores de matriz de entradas.
	var inputsIndex int
	var labelsIndex int

	// Mover secuencialmente las filas en slice de flotantes
	for idx, record := range rawCSVData {

		// Omitir fila del encabezado.
		if idx == 0 {
			continue
		}

		// Recorrer columnas flotantes.
		for i, val := range record {

			// Convertir valor en flotante.
			parsedVal, err := strconv.ParseFloat(val, 64)
			if err != nil {
				log.Fatal(err)
			}

			// Agregar a labelsData si es relevante.
			if i == 4 || i == 5 || i == 6 {
				labelsData[labelsIndex] = parsedVal
				labelsIndex++
				continue
			}

			// Agregar valor flotante a slice de flotantes.
			inputsData[inputsIndex] = parsedVal
			inputsIndex++
		}
	}

	// Formar las matrices
	inputs := mat.NewDense(len(rawCSVData), 4, inputsData)
	labels := mat.NewDense(len(rawCSVData), 3, labelsData)

	// Definir arquitectura de red y parámetros de aprendizaje.
	config := neuralNetConfig{
		inputNeurons:  4,
		outputNeurons: 3,
		hiddenNeurons: 3,
		numEpochs:     5000,
		learningRate:  0.3,
	}

	// Entrenar red neuronal.
	network := newNetwork(config)
	if err := network.train(inputs, labels); err != nil {
		log.Fatal(err)
	}

	// Abrir archivo de conjunto de datos de prueba.
	f, err = os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Crear nuevo lector CSV que lea del archivo abierto.
	reader = csv.NewReader(f)
	reader.FieldsPerRecord = 7

	// Leer en todos los registros CSV de prueba
	rawCSVData, err = reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// inputsData y labelsData almacenarán los valores
	// flotantes que se usarán para formar las matrices.
	inputsData = make([]float64, 4*len(rawCSVData))
	labelsData = make([]float64, 3*len(rawCSVData))

	// inputsIndex rastreará el índice actual
	// de valores de matriz de entradas.
	inputsIndex = 0
	labelsIndex = 0

	// Mover secuencialmente filas en slice de flotantes.
	for idx, record := range rawCSVData {

		// Omitir fila de encabezado
		if idx == 0 {
			continue
		}

		// Recorrer columnas flotantes
		for i, val := range record {

			// Convertir valor en flotante.
			parsedVal, err := strconv.ParseFloat(val, 64)
			if err != nil {
				log.Fatal(err)
			}

			// Agregar a labelsData si es relevante.
			if i == 4 || i == 5 || i == 6 {
				labelsData[labelsIndex] = parsedVal
				labelsIndex++
				continue
			}

			// Agregar valor flotante a slice de flotantes.
			inputsData[inputsIndex] = parsedVal
			inputsIndex++
		}
	}

	// Formar las matrices.
	testInputs := mat.NewDense(len(rawCSVData), 4, inputsData)
	testLabels := mat.NewDense(len(rawCSVData), 3, labelsData)

	// hacer predicciones usando el modelo entrenado.
	predictions, err := network.predict(testInputs)
	if err != nil {
		log.Fatal(err)
	}

	// Calcular la precisión del modelo.
	var truePosNeg int
	numPreds, _ := predictions.Dims()
	for i := 0; i < numPreds; i++ {

		// Otener etiqueta
		labelRow := mat.Row(nil, i, testLabels)
		var species int
		for idx, label := range labelRow {
			if label == 1.0 {
				species = idx
				break
			}
		}

		// Acumular el contador de verdadero positivo/negativo
		if predictions.At(i, species) == floats.Max(mat.Row(nil, i, predictions)) {
			truePosNeg++
		}
	}

	// Calcular precisión (precisión del subconjunto).
	accuracy := float64(truePosNeg) / float64(numPreds)

	// Salida valor de precisión a salida estándar.
	fmt.Printf("\nAccuracy = %0.2f\n\n", accuracy)
}

// NewNetwork inicializa una nueva red neuronal.
func newNetwork(config neuralNetConfig) *neuralNet {
	return &neuralNet{config: config}
}

// Train entrena una red neuronal usando propagación hacia atrás.
func (nn *neuralNet) train(x, y *mat.Dense) error {

	// Inicializar sesgos/pesos.
	randSource := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSource)

	wHiddenRaw := make([]float64, nn.config.hiddenNeurons*nn.config.inputNeurons)
	bHiddenRaw := make([]float64, nn.config.hiddenNeurons)
	wOutRaw := make([]float64, nn.config.outputNeurons*nn.config.hiddenNeurons)
	bOutRaw := make([]float64, nn.config.outputNeurons)

	for _, param := range [][]float64{wHiddenRaw, bHiddenRaw, wOutRaw, bOutRaw} {
		for i := range param {
			param[i] = randGen.Float64()
		}
	}

	wHidden := mat.NewDense(nn.config.inputNeurons, nn.config.hiddenNeurons, wHiddenRaw)
	bHidden := mat.NewDense(1, nn.config.hiddenNeurons, bHiddenRaw)
	wOut := mat.NewDense(nn.config.hiddenNeurons, nn.config.outputNeurons, wOutRaw)
	bOut := mat.NewDense(1, nn.config.outputNeurons, bOutRaw)

	// Definir la salida de la red neuronal.
	//output := mat.NewDense(0, 0, nil)
	var output mat.Dense

	// Recorrer número de épocas que utilizan la
	// propagación hacia atrás para entrenar modelo.
	for i := 0; i < nn.config.numEpochs; i++ {

		// Completar proceso de retroalimentación.
		// hiddenLayerInput := mat.NewDense(0, 0, nil)
		var hiddenLayerInput mat.Dense
		hiddenLayerInput.Mul(x, wHidden)
		addBHidden := func(_, col int, v float64) float64 { return v + bHidden.At(0, col) }
		hiddenLayerInput.Apply(addBHidden, &hiddenLayerInput)

		// hiddenLayerActivations := mat.NewDense(0, 0, nil)
		var hiddenLayerActivations mat.Dense
		applySigmoid := func(_, _ int, v float64) float64 { return sigmoid(v) }
		hiddenLayerActivations.Apply(applySigmoid, &hiddenLayerInput)

		// outputLayerInput := mat.NewDense(0, 0, nil)
		var outputLayerInput mat.Dense
		outputLayerInput.Mul(&hiddenLayerActivations, wOut)
		addBOut := func(_, col int, v float64) float64 { return v + bOut.At(0, col) }
		outputLayerInput.Apply(addBOut, &outputLayerInput)
		output.Apply(applySigmoid, &outputLayerInput)

		// Completar propagación hacia atrás.
		// networkError := mat.NewDense(0, 0, nil)
		var networkError mat.Dense
		networkError.Sub(y, &output)

		// slopeOutputLayer := mat.NewDense(0, 0, nil)
		var slopeOutputLayer mat.Dense
		applySigmoidPrime := func(_, _ int, v float64) float64 { return sigmoidPrime(v) }
		slopeOutputLayer.Apply(applySigmoidPrime, &output)
		// slopeHiddenLayer := mat.NewDense(0, 0, nil)
		var slopeHiddenLayer mat.Dense
		slopeHiddenLayer.Apply(applySigmoidPrime, &hiddenLayerActivations)

		// dOutput := mat.NewDense(0, 0, nil)
		var dOutput mat.Dense
		dOutput.MulElem(&networkError, &slopeOutputLayer)
		// errorAtHiddenLayer := mat.NewDense(0, 0, nil)
		var errorAtHiddenLayer mat.Dense
		errorAtHiddenLayer.Mul(&dOutput, wOut.T())

		// dHiddenLayer := mat.NewDense(0, 0, nil)
		var dHiddenLayer mat.Dense
		dHiddenLayer.MulElem(&errorAtHiddenLayer, &slopeHiddenLayer)

		// Ajustar parámetros.
		// wOutAdj := mat.NewDense(0, 0, nil)
		var wOutAdj mat.Dense
		wOutAdj.Mul(hiddenLayerActivations.T(), &dOutput)
		wOutAdj.Scale(nn.config.learningRate, &wOutAdj)
		wOut.Add(wOut, &wOutAdj)

		bOutAdj, err := sumAlongAxis(0, &dOutput)
		if err != nil {
			return err
		}
		bOutAdj.Scale(nn.config.learningRate, bOutAdj)
		bOut.Add(bOut, bOutAdj)

		// wHiddenAdj := mat.NewDense(0, 0, nil)
		var wHiddenAdj mat.Dense
		wHiddenAdj.Mul(x.T(), &dHiddenLayer)
		wHiddenAdj.Scale(nn.config.learningRate, &wHiddenAdj)
		wHidden.Add(wHidden, &wHiddenAdj)

		bHiddenAdj, err := sumAlongAxis(0, &dHiddenLayer)
		if err != nil {
			return err
		}
		bHiddenAdj.Scale(nn.config.learningRate, bHiddenAdj)
		bHidden.Add(bHidden, bHiddenAdj)
	}

	// Definir red neuronal entrenada.
	nn.wHidden = wHidden
	nn.bHidden = bHidden
	nn.wOut = wOut
	nn.bOut = bOut

	return nil
}

// predict hace predicción basada en una red neuronal entrenada.
func (nn *neuralNet) predict(x *mat.Dense) (*mat.Dense, error) {

	// Verificar que valor neuralNet represente un modelo entrenado.
	if nn.wHidden == nil || nn.wOut == nil || nn.bHidden == nil || nn.bOut == nil {
		return nil, errors.New("the supplied neurnal net weights and biases are empty")
	}

	// Definir salida de red neuronal.
	// output := mat.NewDense(0, 0, nil)
	var output mat.Dense

	// Completar proceso de retroalimentación.
	// hiddenLayerInput := mat.NewDense(0, 0, nil)
	var hiddenLayerInput mat.Dense
	hiddenLayerInput.Mul(x, nn.wHidden)
	addBHidden := func(_, col int, v float64) float64 { return v + nn.bHidden.At(0, col) }
	hiddenLayerInput.Apply(addBHidden, &hiddenLayerInput)

	// hiddenLayerActivations := mat.NewDense(0, 0, nil)
	var hiddenLayerActivations mat.Dense
	applySigmoid := func(_, _ int, v float64) float64 { return sigmoid(v) }
	hiddenLayerActivations.Apply(applySigmoid, &hiddenLayerInput)

	// outputLayerInput := mat.NewDense(0, 0, nil)
	var outputLayerInput mat.Dense
	outputLayerInput.Mul(&hiddenLayerActivations, nn.wOut)
	addBOut := func(_, col int, v float64) float64 { return v + nn.bOut.At(0, col) }
	outputLayerInput.Apply(addBOut, &outputLayerInput)
	output.Apply(applySigmoid, &outputLayerInput)

	return &output, nil
}

// sigmoid implementa función sigmoid
// para uso en funciones de activación.
func sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

// sigmoidPrime implementa la derivada
// de la función sigmoid para la retropropagación.
func sigmoidPrime(x float64) float64 {
	return x * (1.0 - x)
}

// sumAlongAxis suma una matriz a lo largo
// de una dimensión particular, preservando la otra dimensión.
func sumAlongAxis(axis int, m *mat.Dense) (*mat.Dense, error) {

	numRows, numCols := m.Dims()

	var output *mat.Dense

	switch axis {
	case 0:
		data := make([]float64, numCols)
		for i := 0; i < numCols; i++ {
			col := mat.Col(nil, i, m)
			data[i] = floats.Sum(col)
		}
		output = mat.NewDense(1, numCols, data)
	case 1:
		data := make([]float64, numRows)
		for i := 0; i < numRows; i++ {
			row := mat.Row(nil, i, m)
			data[i] = floats.Sum(row)
		}
		output = mat.NewDense(numRows, 1, data)
	default:
		return nil, errors.New("invalid axis, must be 0 or 1")
	}

	return output, nil
}
