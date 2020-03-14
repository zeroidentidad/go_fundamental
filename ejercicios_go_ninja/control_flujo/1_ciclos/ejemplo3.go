package main

import (
	"fmt"
)

func main() {
	// for init; condition; post {}
	for i := 0; i <= 10; i++ {
		fmt.Printf("El ciclo externo: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t El ciclo interno: %d\n", j)

		}
	}
}
