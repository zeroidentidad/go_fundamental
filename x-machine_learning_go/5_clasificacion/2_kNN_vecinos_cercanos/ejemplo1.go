package main

import (
	"fmt"
	"log"
	"math"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

func main() {

	// Leer en el conjunto de datos de iris en golearn "instances".
	irisData, err := base.ParseCSVToInstances("./data/iris.csv", true)
	if err != nil {
		log.Fatal(err)
	}

	// Inicializar un nuevo clasificador KNN. Usando una simple
	// medida de distancia Euclidiana yk = 2.
	knn := knn.NewKnnClassifier("euclidean", "linear", 2)

	// Usar la validación cruzada para entrenar y evaluar sucesivamente
	// el modelo en 5 pliegues del conjunto de datos.
	cv, err := evaluation.GenerateCrossFoldValidationConfusionMatrices(irisData, knn, 5)
	if err != nil {
		log.Fatal(err)
	}

	// Obtener la media, la varianza y la desviación estándar de la precisión para la validación cruzada.
	mean, variance := evaluation.GetCrossValidatedMetric(cv, evaluation.GetAccuracy)
	stdev := math.Sqrt(variance)

	// Salida de las métricas cruzadas a la salida estándar.
	fmt.Printf("\nAccuracy\n%.2f (+/- %.2f)\n\n", mean, stdev*2)
}
