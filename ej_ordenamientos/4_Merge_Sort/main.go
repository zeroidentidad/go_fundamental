package main

import (
	"fmt"
)

func main() {
	arr_original := []int{31, 41, 59, 26, 53, 58, 97, 93, 23, 84, 1, 100, 5, 23}
	fmt.Printf("Resultado:\n%v\n", merge_sort(arr_original))
}

func merge_sort(arreglo []int) []int {
	// evaluar si arreglo recibido tiene 1 elemento
	if len(arreglo) == 1 {
		return arreglo
	}
	// partir arreglo
	arreglo_size := len(arreglo)
	// ordenar partidos
	mitad := arreglo_size / 2
	izquierda := arreglo[0:mitad]
	derecha := arreglo[mitad:arreglo_size]
	fmt.Println(izquierda)
	fmt.Println(derecha)

	// ordenar partidos
	izquierda = merge_sort(izquierda)
	derecha = merge_sort(derecha)

	// combinar

	return merge(izquierda, derecha)
}

func merge(izquierda, derecha []int) []int {
	// si ultimo izquierda menor que primer derecha
	ultimo_izquierdo := izquierda[len(izquierda)-1]
	primer_derecho := derecha[0]
	if ultimo_izquierdo <= primer_derecho {
		return append(izquierda, derecha...)
	}

	// nuevo arreglo
	resultado := make([]int, 0) // creacion en slice

	var menor int

	for len(izquierda) > 0 && len(derecha) > 0 {
		// combinar
		if izquierda[0] <= derecha[0] {
			menor = izquierda[0]
			izquierda = izquierda[1:]
		} else {
			menor = derecha[0]
			derecha = derecha[1:]
		}
		// elemento en nuevo arreglo
		resultado = append(resultado, menor)
	}

	// con elementos en arreglos, sumar
	if len(izquierda) > 0 {
		resultado = append(resultado, izquierda...)
	}
	if len(derecha) > 0 {
		resultado = append(resultado, derecha...)
	}

	return resultado
}
