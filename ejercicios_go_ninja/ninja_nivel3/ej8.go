package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("No debería imprimir")
	case true:
		fmt.Println("Debería imprimir")
	}
}
