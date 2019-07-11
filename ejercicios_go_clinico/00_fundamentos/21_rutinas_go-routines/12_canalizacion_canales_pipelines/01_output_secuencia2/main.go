package main

import "fmt"

func main() {
	// Set up el pipeline y consume la salida.
	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n) // 16 entonces 81
	}
}

func gen(nums ...int) chan int {
	salida := make(chan int)
	go func() {
		for _, n := range nums {
			salida <- n
		}
		close(salida)
	}()
	return salida
}

func sq(in chan int) chan int {
	salida := make(chan int)
	go func() {
		for n := range in {
			salida <- n * n
		}
		close(salida)
	}()
	return salida
}
