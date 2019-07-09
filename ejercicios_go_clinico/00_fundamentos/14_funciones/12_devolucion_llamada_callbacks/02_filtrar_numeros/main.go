package main

import "fmt"

func filtro(numeros []int, callback func(int) bool) []int {
	var xs []int
	for _, n := range numeros {
		if callback(n) {
			xs = append(xs, n) // agregar elemento
		}
	}
	return xs
}

func main() {
	xs := filtro([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})
	fmt.Println(xs) // [2 3 4]
}
