package main

import (
	"fmt"
)

func main() {
	arr_original := []int{31, 41, 59, 26, 53, 58, 97, 93, 23, 84, 40, 1, 150}
	quick_sort(&arr_original, 0, len(arr_original)-1)
	fmt.Println(arr_original)
}

func quick_sort(arreglo *[]int, izquierda, derecha int) []int {
	// separar valores menores a izquierda de pivot
	// y valores mayores a la derecha del pivot
	if izquierda < derecha {
		arr := *arreglo
		limite, origen := derecha, izquierda
		pivot := arr[derecha]
		derecha--

		for izquierda <= derecha {
			// buscar en izquierda numero mayor que pivot
			for izquierda < len(arr) && arr[izquierda] < pivot {
				izquierda++
			}
			// buscar en derecha numero menor que pivot
			for derecha >= 0 && arr[derecha] > pivot {
				derecha--
			}

			if izquierda <= derecha {
				// intercambiar encontrados
				swap(arreglo, izquierda, derecha)
				// aumentar valores izquierda
				izquierda++
				// decrementar valores derecha
				derecha--
			}

		}
		// termina separacion izquierda | derecha
		swap(arreglo, izquierda, limite)
		quick_sort(arreglo, origen, derecha)
		quick_sort(arreglo, izquierda, limite)
	}

	return *arreglo
}

func swap(arreglo *[]int, izquierda, derecha int) {
	// intercambiar valor izquierda con derecha
	arr := *arreglo
	temporal := arr[izquierda]
	arr[izquierda] = arr[derecha]
	arr[derecha] = temporal
}
