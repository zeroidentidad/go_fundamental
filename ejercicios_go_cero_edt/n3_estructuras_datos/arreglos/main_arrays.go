package main

import "fmt"

func main() {
	// Arrays
	// Todos los valores deben ser del mismo tipo de dato
	// Tamaño fijo.

	//var nombres [3]string
	/*
		nombres[0] = "Jesus"
		nombres[1] = "Karent"
		nombres[2] = "Vero"
		// nombres[3] = "Pedro"
	*/
	nombres := [3]string{
		"Jesus",
		"Karent",
		"Vero",
	}
	//nombre := nombres[1]
	size := len(nombres)

	fmt.Println("Tamaño arreglo de:", size)
	fmt.Println(nombres)

	// fmt.Println(nombres[inicio:final(excluyente)])
	// fmt.Println(nombres[1:2])
}
