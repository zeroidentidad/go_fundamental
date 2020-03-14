package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)

	y := append(x, 6, 7, 8) //Un nuevo arreglo subyacente es asignado
	fmt.Println(x)
	fmt.Println(y)
}
