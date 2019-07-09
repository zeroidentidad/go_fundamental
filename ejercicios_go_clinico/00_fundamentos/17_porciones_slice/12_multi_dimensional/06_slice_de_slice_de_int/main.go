package main

import (
	"fmt"
)

func main() {

	transacciones := make([][]int, 0, 3)

	for i := 0; i < 3; i++ {
		transaccion := make([]int, 0, 4)
		for j := 0; j < 4; j++ {
			transaccion = append(transaccion, j)
		}
		transacciones = append(transacciones, transaccion)
	}
	fmt.Println(transacciones)
}
