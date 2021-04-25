package main

import (
	"fmt"
	"math/cmplx"
)

func main() {

	// los números complejos son
	// definidos como reales e imaginarios
	// parte definida por float64
	a := complex(2, 3)

	fmt.Printf("Parte real: %f \n", real(a))
	fmt.Printf("Parte compleja: %f \n", imag(a))

	b := complex(6, 4)

	// Todos los operadores comunes son útiles.
	c := a - b
	fmt.Printf("Diferencia : %v\n", c)
	c = a + b
	fmt.Printf("Suma : %v\n", c)
	c = a * b
	fmt.Printf("Producto : %v\n", c)
	c = a / b
	fmt.Printf("Division : %v\n", c)

	conjugate := cmplx.Conj(a)
	fmt.Println("Conjugado número complejo de variable a : ", conjugate)

	cos := cmplx.Cos(b)
	fmt.Println("Coseno de b : ", cos)

}
