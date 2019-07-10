package main

import "fmt"

func main() {
	residuo := 7.24
	fmt.Printf("%T\n", residuo)
	fmt.Printf("%T\n", int(residuo))

	var valor interface{} = 7
	fmt.Printf("%T\n", valor)
	fmt.Printf("%T\n", int(valor))
	// cannot convert valor (type interface {}) to type int: need type assertion
	// correcto: fmt.Printf("%T\n", valor.(int))
}
