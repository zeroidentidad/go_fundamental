package main

import "fmt"

func main() {

	/*
		var n1, n2 int8
		n1 = 15
		n2 = 3
		r := suma(n1, n2)
		fmt.Println(r)
	*/

	var edad uint8
	edad = 30
	fmt.Println(tipoEdad(edad))
}

func suma(a, b int8) int8 {
	return a + b
}

func tipoEdad(edad uint8) string {
	var tipo string
	switch {
	case edad < 12:
		tipo = "Niñez"
	case edad < 18:
		tipo = "Adolescencia"
	default:
		tipo = "Madurez"
	}
	return tipo
}
