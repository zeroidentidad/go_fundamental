package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {

	// Crear una representación plana de la matriz.
	data := []float64{1.2, -5.7, -2.4, 7.3}

	// Formar la matriz.
	a := mat.NewDense(2, 2, data)

	// Obtener un solo valor de la matriz.
	val := a.At(0, 1)
	fmt.Printf("El valor de a en (0,1) es: %.2f\n\n", val)

	// Get the values in a specific column.
	col := mat.Col(nil, 0, a)
	fmt.Printf("El valor de a en (0,1) es: %v\n\n", col)

	// Obtener los valores en una fila específica.
	row := mat.Row(nil, 1, a)
	fmt.Printf("Los valores en la segunda fila son: %v\n\n", row)

	// Modificar un solo elemento.
	a.Set(0, 1, 11.2)

	// Modificar una fila completa.
	a.SetRow(0, []float64{14.3, -4.2})

	// Modificar una columna completa.
	a.SetCol(0, []float64{1.7, -0.3})

	// Verificación de integridad, enviar la matriz a la salida estándar.
	fa := mat.Formatted(a, mat.Prefix("    "))
	fmt.Printf("A = %v\n\n", fa)
}
