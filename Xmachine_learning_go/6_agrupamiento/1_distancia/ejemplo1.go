package main

import (
	"fmt"

	"gonum.org/v1/gonum/floats"
)

func main() {

	// Calcular la distancia Euclidiana, especificada
	// a través del último argumento en la función Distance.
	distance := floats.Distance([]float64{1, 2}, []float64{3, 4}, 2)

	fmt.Printf("\nDistance: %0.2f\n\n", distance)
}
