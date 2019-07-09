package main

import "fmt"

func main() {
	var x [58]int // 58 espacios
	fmt.Println(x)
	fmt.Println(len(x)) // longitud de 0 - 57
	fmt.Println(x[42])  // 0 - sin valor asignado
	x[42] = 777
	fmt.Println(x[42]) // 777 - valor asignado
}
