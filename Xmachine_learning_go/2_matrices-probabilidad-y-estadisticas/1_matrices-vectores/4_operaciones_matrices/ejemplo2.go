package main

import (
	"fmt"
	"log"

	"gonum.org/v1/gonum/mat"
)

func main() {

	// Crear una nueva matriz a.
	a := mat.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})

	// Calcular y generar la transposición de la matriz.
	ft := mat.Formatted(a.T(), mat.Prefix("      "))
	fmt.Printf("a^T = %v\n\n", ft)

	// Calcular y generar el determinante de a.
	deta := mat.Det(a)
	fmt.Printf("det(a) = %.2f\n\n", deta)

	// Calcular y generar el inverso de a.
	//aInverse := mat.NewDense(0, 0, nil) // <- panic
	//aInverse := mat.NewDense(3, 3, nil) // <- salida extra
	var aInverse mat.Dense // fixed mismo resultado
	if err := aInverse.Inverse(a); err != nil {
		log.Fatal(err)
	}
	fi := mat.Formatted(&aInverse, mat.Prefix("       ")) // mod * -> &
	fmt.Printf("a^-1 = %v\n\n", fi)
}
