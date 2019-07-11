package main

import (
	"fmt"
)

func main() {
	c := factorial(4)
	for n := range c {
		fmt.Println(n)
	}
}

func factorial(n int) chan int {
	salida := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		salida <- total
		close(salida)
	}()
	return salida
}

// para un cálculo como este con una goroutine lleva más tiempo en realizarse por instanciar
// mas componentes del lenguaje para un proceso lineal
