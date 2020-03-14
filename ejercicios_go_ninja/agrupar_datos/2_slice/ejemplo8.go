package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)

	y := append(x[:2], x[3:]...) //Se utiliza el mismo arreglo subyacente
	fmt.Println(x)
	fmt.Println(y)
}
