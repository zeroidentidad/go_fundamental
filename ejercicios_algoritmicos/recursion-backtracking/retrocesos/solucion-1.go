package main

import "fmt"

func main() {
	A := []int{2, 45, 67, 8, 4}
	generateBinaryStrings(5, A, 3)
}

func printResult(A []int, n int) {
	var i int
	for ; i < n; i++ {
		// Función para imprimir la salida
		fmt.Print(A[i])
	}
	fmt.Printf("\n")
}

// Función para generar todas las cadenas binarias
func generateBinaryStrings(n int, A []int, i int) {
	// n total a usar, i elementos restantes
	if i == n {
		printResult(A, n)
		return
	}

	// asignar "0" en la posición i-ésima y probar todas las demás permutaciones para las posiciones restantes
	A[i] = 0
	generateBinaryStrings(n, A, i+1)

	// asignar "1" en la posición i-ésima e intentar todas las demás permutaciones para las posiciones restantes
	A[i] = 1
	generateBinaryStrings(n, A, i+1)
}
