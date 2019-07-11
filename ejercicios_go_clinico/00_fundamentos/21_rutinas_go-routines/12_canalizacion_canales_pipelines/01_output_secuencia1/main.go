package main

import "fmt"

func main() {
	// Set up del pipeline.
	c := gen(2, 3)
	salida := sq(c)

	// Consume la salida.
	fmt.Println(<-salida) // 4
	fmt.Println(<-salida) // 9
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
