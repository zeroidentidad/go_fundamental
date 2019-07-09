package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)
	// 3 es la longitud - número de elementos referidos por el slice
	// 5 es capacidad - número de elementos en la matriz subyacente

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"
	greeting[3] = "suprabadham" // fuera de la logitud indicada

	fmt.Println(greeting[2])
}
