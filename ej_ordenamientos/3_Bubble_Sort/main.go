package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	arr_original := []int{412, 12, 123, 1, 2324, 3, 155, 23412, 2}

	start := time.Now()
	arr_ordenado := bubble_sort(arr_original)

	fmt.Println(arr_ordenado)
	elapsed := time.Since(start)
	log.Printf("\nBubble Sort tomó: %s", elapsed)
}

func bubble_sort(arreglo []int) []int {
	numeros_ordenados := 0
	//numero_iteraciones := 0
	intercambio := true
	// mientras hay intercambios
	for intercambio {
		// iterar arreglo
		intercambio = false
		for i := 1; i < (len(arreglo) - numeros_ordenados); i++ { // sin optimizacion len(arreglo)
			//numero_iteraciones++
			// comparar pares
			if arreglo[i-1] > arreglo[i] {
				// si izquierdo mayor que derecho, intercambiar
				intercambio = true
				swap(&arreglo, i)
			}
		}
		numeros_ordenados++
	}

	// salida iteraciones necesarias
	// fmt.Println(numero_iteraciones)
	return arreglo
}

func swap(puntero_arreglo *[]int, index_derecha int) {
	index_izquierda := index_derecha - 1
	arreglo := *puntero_arreglo
	copia := arreglo[index_izquierda]
	arreglo[index_izquierda] = arreglo[index_derecha]
	arreglo[index_derecha] = copia

}
