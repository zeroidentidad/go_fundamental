package main

import (
	"fmt"
)

func main() {

	in := gen()

	f := factorial(in)

	for n := range f {
		fmt.Println(n)
	}
}

func gen() <-chan int {
	salida := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				salida <- j
			}
		}
		close(salida)
	}()
	return salida
}

func factorial(in <-chan int) <-chan int {
	salida := make(chan int)
	go func() {
		for n := range in {
			salida <- fact(n)
		}
		close(salida)
	}()
	return salida
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
