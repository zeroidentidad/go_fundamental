package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // 0xc0000180f0

	var b = &a
	fmt.Println(b)  // 0xc0000180f0
	fmt.Println(*b) // 43

	// b es un puntero int;
	// b apunta a la dirección de memoria donde se almacena un int
	// para ver el valor en esa dirección de memoria, agregue un * delante de b
	// esto se conoce como desreferenciación
	// el * es un operador en este caso
}
