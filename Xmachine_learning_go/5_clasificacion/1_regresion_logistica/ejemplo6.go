package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"

	"gonum.org/v1/gonum/mat" // "github.com/gonum/matrix/mat64"
)

func main() {

	// Parte inicial entrenamiento y prueba de regresion logistica

	// Abrir archivo de conjunto de datos de entrenamiento.
	f, err := os.Open("./newData/training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Crear un nuevo lector CSV que lea del archivo abierto.
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 2

	// Leer en todos los registros CSV
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// featureData y las etiquetas contendrán todos los valores flotantes
	// que eventualmente se utilizarán en entrenamiento.
	featureData := make([]float64, 2*(len(rawCSVData)-1))
	labels := make([]float64, len(rawCSVData)-1)

	// featureIndex rastreará el índice actual de los valores de la matriz de características.
	var featureIndex int

	// Mover secuencialmente las filas en el slice de floats.
	for idx, record := range rawCSVData {

		// Omitir fila de encabezado.
		if idx == 0 {
			continue
		}

		// Agregar la caracteristica de puntaje FICO.
		featureVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		featureData[featureIndex] = featureVal

		// Agregar una intercepción.
		featureData[featureIndex+1] = 1.0

		// Incrementar fila de características.
		featureIndex += 2

		// Agregar la etiqueta de clase.
		labelVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		labels[idx-1] = labelVal
	}

	// Formar matriz a partir de las características.
	features := mat.NewDense(len(rawCSVData)-1, 2, featureData)

	// Entrenar el modelo de regresión logística.
	weights := logisticRegression(features, labels, 1000, 0.3)

	// Enviar la fórmula del modelo de Regresión logística a stdout.
	formula := "p = 1 / ( 1 + exp(- m1 * FICO.score - m2) )"
	fmt.Printf("\n%s\n\nm1 = %0.2f\nm2 = %0.2f\n\n", formula, weights[0], weights[1])
}

// Logistic implementa la función logistic,
// que se utiliza en la regresión logística.
func logistic(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

// logisticRegression se ajusta a un modelo de regresión logística para los datos dados.
func logisticRegression(features *mat.Dense, labels []float64, numSteps int, learningRate float64) []float64 {

	// Inicializar pesos aleatorios.
	_, numWeights := features.Dims()
	weights := make([]float64, numWeights)

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for idx, _ := range weights {
		weights[idx] = r.Float64()
	}

	/// Optimizar iterativamente los pesos.
	for i := 0; i < numSteps; i++ {

		// Inicializar una variable para acumular errores para esta iteración.
		var sumError float64

		// Hacer predicciones para cada etiqueta y acumular errores.
		for idx, label := range labels {

			// Obtener las características correspondientes a la etiqueta.
			featureRow := mat.Row(nil, idx, features)

			// Calcular el error para los pesos de la iteración.
			pred := logistic(featureRow[0]*weights[0] + featureRow[1]*weights[1])
			predError := label - pred
			sumError += math.Pow(predError, 2)

			// Actualizar pesos de las características.
			for j := 0; j < len(featureRow); j++ {
				weights[j] += learningRate * predError * pred * (1 - pred) * featureRow[j]
			}
		}
	}

	return weights
}
