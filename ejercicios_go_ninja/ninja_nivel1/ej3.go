package main

import (
	"fmt"
)

type numero int

var x numero

func main() {
	fmt.Println(x)
	fmt.Printf("El tipo de x es: %T\n", x)

	x = 42
	fmt.Println(x)
}
