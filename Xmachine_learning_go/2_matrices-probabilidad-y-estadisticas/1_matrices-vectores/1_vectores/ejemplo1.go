package main

import "fmt"

func main() {

	// Inicializar un "vector" a través de un slice.
	var myvector []float64

	// Agregar un par de componentes al vector.
	myvector = append(myvector, 11.0)
	myvector = append(myvector, 5.2)

	fmt.Println(myvector)
}
