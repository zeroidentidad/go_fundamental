package main

import (
	"fmt"

	"gonum.org/v1/gonum/floats"
)

func main() {

	// Inicializar un par de "vectores" representados como slices.
	vectorA := []float64{11.0, 5.2, -1.3}
	vectorB := []float64{-7.2, 4.2, 5.1}

	// Calcular el producto escalar de A y B
	// (https://en.wikipedia.org/wiki/Dot_product).
	dotProduct := floats.Dot(vectorA, vectorB)
	fmt.Printf("El producto escalar de A y B es: %0.2f\n", dotProduct)

	// Escalar cada elemento de A en 1.5.
	floats.Scale(1.5, vectorA)
	fmt.Printf("Escalando A por 1.5 da: %v\n", vectorA)

	// Calcular la norma/longitud de B.
	normB := floats.Norm(vectorB, 2)
	fmt.Printf("La norma/longitud de B es: %0.2f\n", normB)
}
