package main

import "fmt"

func main() {

	var resultados []int
	fmt.Println(resultados)

	miSlice := []string{"a", "b", "c", "g", "m", "z"}
	fmt.Println(miSlice)
	fmt.Println(miSlice[2:4])  // partir slice
	fmt.Println(miSlice[2])    // acceso al índice; acceso por índice
	fmt.Println("miString"[2]) // acceso al índice; acceso por índice
}
