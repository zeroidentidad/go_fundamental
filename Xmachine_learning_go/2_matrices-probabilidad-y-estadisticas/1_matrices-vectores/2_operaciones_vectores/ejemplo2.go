package main

import (
	"fmt"

	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/mat"
)

func main() {

	// Inicializar un par de "vectores" representados como slices.
	vectorA := mat.NewVecDense(3, []float64{11.0, 5.2, -1.3})
	vectorB := mat.NewVecDense(3, []float64{-7.2, 4.2, 5.1})

	// Calcular el producto escalar de A y B
	// (https://en.wikipedia.org/wiki/Dot_product).
	dotProduct := mat.Dot(vectorA, vectorB)
	fmt.Printf("El producto escalar de A y B es: %0.2f\n", dotProduct)

	// Escalar cada elemento de A en 1.5
	vectorA.ScaleVec(1.5, vectorA)
	fmt.Printf("Escalando A por 1.5 da: %v\n", vectorA)

	// Compute the norm/length of B.
	normB := blas64.Nrm2(vectorB.RawVector())
	fmt.Printf("La norma/longitud de B es: %0.2f\n", normB)
}
