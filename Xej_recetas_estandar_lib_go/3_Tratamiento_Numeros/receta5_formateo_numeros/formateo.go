package main

import (
	"fmt"
)

var integer int64 = 32500
var floatNum float64 = 22000.456

func main() {

	// Forma común de cómo imprimir a decimal
	fmt.Printf("%d \n", integer)

	// Para siempre mostrar el signo
	fmt.Printf("%+d \n", integer)

	// Imprimir en otra base x -16, o -8, b -2, d -10
	fmt.Printf("%x \n", integer)
	fmt.Printf("%#x \n", integer)

	// Relleno con ceros a la izquierda
	fmt.Printf("%010d \n", integer)

	// Relleno con espacios a la izquierda
	fmt.Printf("% 10d \n", integer)

	// Padding derecho
	fmt.Printf("% -10d \n", integer)

	// Imprimir número de punto flotante
	fmt.Printf("%f \n", floatNum)

	// Número de punto flotante con precisión limitada a 5
	fmt.Printf("%.5f \n", floatNum)

	// Número de punto flotante en notación científica
	fmt.Printf("%e \n", floatNum)

	// Número de punto flotante %e para exponentes grandes
	// o %f de lo contrario
	fmt.Printf("%g \n", floatNum)

}
