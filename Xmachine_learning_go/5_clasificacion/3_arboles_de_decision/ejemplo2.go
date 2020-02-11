package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/ensemble"
	"github.com/sjwhitworth/golearn/evaluation"
)

func main() {

	// Cambio a Random Forest (bosque aleatorio)

	// Leer el conjunto de datos iris en golearn "instances".
	irisData, err := base.ParseCSVToInstances("./data/iris.csv", true)
	if err != nil {
		log.Fatal(err)
	}

	// Esto es para sembrar procesos aleatorios involucrados
	// en la construcción del árbol de decisión.
	rand.Seed(44111342)

	// Armar bosque aleatorio con 10 árboles y 2 características por árbol,
	// lo cual es un valor predeterminado sensato (el número de características
	// por árbol normalmente se establece en sqrt (número de características)).
	rf := ensemble.NewRandomForest(10, 2) // con 4 similar a uso datos arbol simple

	// Usar la validación cruzada para entrenar y evaluar sucesivamente
	// el modelo en 5 pliegues del conjunto de datos.
	cv, err := evaluation.GenerateCrossFoldValidationConfusionMatrices(irisData, rf, 5)
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
