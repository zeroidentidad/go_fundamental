package main

import "fmt"

func main() {
	opciones := []string{"A1", "B2", "C3"}
	perm := permutaciones(opciones)
	for _, v := range perm {
		fmt.Println(v)
	}
	fmt.Println("Numero de respuestas:", len(perm))
}

func permutaciones(opciones []string) [][]string {
	return permutar(opciones, 0)
}

func swap(a []string, x, y int) {
	a[x], a[y] = a[y], a[x]
}

func permutar(opciones []string, start int) [][]string {
	respuestas := [][]string{}
	end := len(opciones) - 1
	if start == end {
		resp := make([]string, len(opciones))
		copy(resp, opciones)
		respuestas = append(respuestas, resp)
	} else {
		for i := start; i <= end; i++ {
			swap(opciones, start, i)
			subRespuestas := permutar(opciones, start+1)
			for _, resp := range subRespuestas {
				respuestas = append(respuestas, resp)
			}
			swap(opciones, start, i)
		}
	}
	return respuestas
}
