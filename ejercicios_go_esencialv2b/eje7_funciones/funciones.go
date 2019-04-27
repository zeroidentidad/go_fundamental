package main

import (
	"fmt"
)

func main() {
	fmt.Println(multiplicar(10, 12), multiplicar(4, 5))
}

func multiplicar(num1, num2 int) int {
	return num1 * num2
}
