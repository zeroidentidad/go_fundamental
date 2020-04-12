package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	arr_original := []int{412, 12, 123, 1, 2324, 3, 155, 23412, 2}

	start := time.Now()
	arr_ordenado := selection_sort(arr_original)

	fmt.Println(arr_ordenado)

	elapsed := time.Since(start)
	log.Printf("\nSelection Sort tomó: %s", elapsed)
}

func selection_sort(arreglo []int) []int {
	for i := 0; i < len(arreglo); i++ {
		minimo_encontrado, posicion_minimo := arreglo[i], i

		valor_original := arreglo[i]
		// encontrar minimo en parte desordenada
		for j := i + 1; j < len(arreglo); j++ {
			valor_comparacion := arreglo[j]
			if valor_comparacion < minimo_encontrado {
				minimo_encontrado, posicion_minimo = valor_comparacion, j
			}
		}

		if minimo_encontrado != valor_original {
			// intercambio posiciones con primer desordenado
			arreglo[i], arreglo[posicion_minimo] = minimo_encontrado, valor_original
		}
	}

	return arreglo
}
