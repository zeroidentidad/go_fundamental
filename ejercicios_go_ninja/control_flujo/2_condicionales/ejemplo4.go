package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("Esta no se debe imprimir")
	case (2 == 4):
		fmt.Println("Esta no se debe imprimir2")
	case (3 == 3):
		fmt.Println("Imprime")
		fallthrough
	case (4 == 4):
		fmt.Println("también verdadera, se imprime?")
		fallthrough
	case (7 == 9):
		fmt.Println("no verdadero1")
		fallthrough
	case (11 == 14):
		fmt.Println("no verdadero 2")
		fallthrough
	case (15 == 15):
		fmt.Println("verdadero 15")
		fallthrough
	default:
		fmt.Println("Este es default")
	}
}
