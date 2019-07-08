package main

import (
	"fmt"
)

func main() {

	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b = &a

	fmt.Println(b)

	// el código anterior hace que b sea un puntero a la dirección de memoria donde se almacena un int
	// b es de tipo "int puntero"
	// *int -- el * es parte del tipo -- b es del tipo * int
}
