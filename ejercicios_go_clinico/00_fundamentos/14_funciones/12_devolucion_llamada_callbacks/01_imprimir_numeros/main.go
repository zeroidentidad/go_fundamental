package main

import "fmt"

func visita(numeros []int, callback func(int)) {
	for _, n := range numeros {
		callback(n)
	}
}

func main() {
	visita([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})
}

// callback: pasando una func como argumento
