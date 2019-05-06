package main

import "fmt"

func main() {
	// Slices
	// 1. Apuntador a un array
	// 2. Tamaño (no es fijo)
	// 3. Capacidad
	// var nombres []string // forma de declaracion simple
	// make (tipo de dato del slice, tamaño inicial, capacidad inicial)
	// nombres := make([]string, 0) // forma de declaracion en la documentacion

	nombres := []string{ // forma de declaración shortcut con valores por default
		"Karent",
		"Vero",
		"Alexandra",
	}

	fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Karla")

	fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Perla")

	fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Guadalupe")

	fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Rubi")

	fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Melva")

	fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
	nombres[3] = "Juan"

	fmt.Println(nombres)
}
