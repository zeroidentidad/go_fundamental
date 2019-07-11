package main

import "fmt"

func main() {
	c := incrementor()
	cSum := extractor(c)
	for n := range cSum {
		fmt.Println(n)
	}
}

func incrementor() <-chan int {
	salida := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			salida <- i
		}
		close(salida)
	}()
	return salida
}

func extractor(c <-chan int) <-chan int {
	salida := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		salida <- sum
		close(salida)
	}()
	return salida
}

/*
El operador opcional <- especifica la dirección del canal, enviar o recibir.
Si no se da una dirección, el canal es bidireccional.
https://golang.org/ref/spec#Channel_types
*/
