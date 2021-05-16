package main

import "fmt"

func main() {
	A := []int{2, 45, 67, 8, 4}
	generateK_aryStrings(5, A, 3, 3)
}

func printResult(A []int, n int) {
	var i int
	for ; i < n; i++ {
		// Función para imprimir la salida
		fmt.Print(A[i])
	}
	fmt.Printf("\n")
}

// Función para generar todas las cadenas k-arias
func generateK_aryStrings(n int, A []int, i int, k int) {
	if i == n {
		printResult(A, n)
		return
	}

	for j := 0; j < k; j++ {
		// Asignar j en la posición i-ésima e intentar todas
		// las demás permutaciones para las posiciones restantes
		A[i] = j
		generateK_aryStrings(n, A, i+1, k)
	}
}
