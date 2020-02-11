package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/trees"
)

func main() {

	// Leer el conjunto de datos iris en golearn "instances".
	irisData, err := base.ParseCSVToInstances("./data/iris.csv", true)
	if err != nil {
		log.Fatal(err)
	}

	// Esto es para sembrar procesos aleatorios involucrados
	// en la construcción del árbol de decisión.
	rand.Seed(44111342)

	// Se usara el algoritmo ID3 para construir el árbol de decisión.
	// Además, se iniciara con un parámetro de 0.6 que controla la división train-prune (entrenar-cortar).
	tree := trees.NewID3DecisionTree(0.6)

	// Usar la validación cruzada para entrenar y evaluar sucesivamente
	// el modelo en 5 pliegues del conjunto de datos.
	cv, err := evaluation.GenerateCrossFoldValidationConfusionMatrices(irisData, tree, 5)
	if err != nil {
		log.Fatal(err)
	}

	// Obtener la media, la varianza y la desviación estándar
	// de la precisión para la validación cruzada.
	mean, variance := evaluation.GetCrossValidatedMetric(cv, evaluation.GetAccuracy)
	stdev := math.Sqrt(variance)

	// Salida de la métrica cruzada a salida estándar.
	fmt.Printf("\nAccuracy\n%.2f (+/- %.2f)\n\n", mean, stdev*2)
}
