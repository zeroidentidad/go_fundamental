package main

import (
	"fmt"
)

func main() {
	// for init; condition; post {}
	for i := 0; i <= 10; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("El cilo externo: %d\t El ciclo interno: %d\n", i, j)

		}
	}
}
