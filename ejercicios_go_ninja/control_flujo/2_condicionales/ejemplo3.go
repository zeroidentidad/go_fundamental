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
	case (4 == 4):
		fmt.Println("también verdadera, se imprime?")
	}
}
