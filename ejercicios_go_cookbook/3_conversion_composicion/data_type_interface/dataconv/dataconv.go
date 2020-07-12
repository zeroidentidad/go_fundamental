package dataconv

import "fmt"

// ShowConv demuestra algún tipo de conversión
func ShowConv() {
	// int
	var a = 24

	// float 64
	var b = 2.0

	// convertir el int a float64 para este cálculo
	c := float64(a) * b
	fmt.Println(c)

	// fmt.Sprintf es una forma simple de convertir a cadenas
	precision := fmt.Sprintf("%.2f", b)

	// imprimir el valor y el tipo
	fmt.Printf("%s - %T\n", precision, precision)
}
