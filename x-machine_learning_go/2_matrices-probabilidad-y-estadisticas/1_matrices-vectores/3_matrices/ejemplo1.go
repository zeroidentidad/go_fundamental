package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {

	// Crear una representación plana de la matriz.
	data := []float64{1.2, -5.7, -2.4, 7.3}

	// Formar matriz.
	a := mat.NewDense(2, 2, data)

	// Verificación de integridad, enviar la matriz a la salida estándar.
	fa := mat.Formatted(a, mat.Prefix("    "))
	fmt.Printf("A = %v\n\n", fa)
}
