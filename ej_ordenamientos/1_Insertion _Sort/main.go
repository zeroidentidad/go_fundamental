package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	arr_original := []int{412, 12, 123, 1, 2324, 3, 155, 23412, 2}

	start := time.Now()
	arr_ordenado := insertion_sort(arr_original)

	fmt.Println("\nResultado:")
	fmt.Println(arr_ordenado)

	elapsed := time.Since(start)
	log.Printf("\nInsertion Sort tomó: %s", elapsed)
}

func insertion_sort(arreglo []int) []int {
	for i := 1; i < len(arreglo); i++ {
		j := i
		for j > 0 && arreglo[j-1] > arreglo[j] {
			// salida intercambios
			//fmt.Printf("%d intercambio > %d\n", arreglo[j-1], arreglo[j])
			swap(j-1, j, &arreglo)
			j--
		}
	}
	return arreglo
}

func swap(previo, actual int, puntero_arreglo *[]int) {
	arreglo := *puntero_arreglo
	copia := arreglo[actual]
	arreglo[actual] = arreglo[previo]
	arreglo[previo] = copia
}
