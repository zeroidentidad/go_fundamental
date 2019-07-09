package main

import "fmt"

func main() {

	saludar := make([]string, 3, 5)

	// 3 es la longitud - número de elementos referidos por el slice
	// 5 es capacidad - número de elementos en la matriz subyacente

	saludar[0] = "Good morning!"
	saludar[1] = "Bonjour!"
	saludar[2] = "buenos dias!"
	saludar = append(saludar, "Suprabadham") // <-

	fmt.Println(saludar[3])
}
