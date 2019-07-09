package main

import "fmt"

func main() {

	numeroCliente := make([]int, 3)
	// 3 es longitud y capacidad
	// // longitud: número de elementos a los que hace referencia el slice
	// // capacidad - número de elementos en la matriz subyacente
	numeroCliente[0] = 7
	numeroCliente[1] = 10
	numeroCliente[2] = 15

	fmt.Println(numeroCliente[0])
	fmt.Println(numeroCliente[1])
	fmt.Println(numeroCliente[2])

	saludar := make([]string, 3, 5)
	// 3 es la longitud - número de elementos referidos en el slice
	// 5 es capacidad - número de elementos en la matriz subyacente
	// también podrías hacerse así

	saludar[0] = "Good morning!"
	saludar[1] = "Bonjour!"
	saludar[2] = "BUenos dias!"

	fmt.Println(saludar[2])
}
