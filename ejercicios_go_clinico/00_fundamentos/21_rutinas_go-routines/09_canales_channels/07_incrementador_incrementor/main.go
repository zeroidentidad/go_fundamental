package main

import (
	"fmt"
)

func main() {
	c1 := incrementor("Foo:")
	c2 := incrementor("Bar:")
	c3 := extractor(c1)
	c4 := extractor(c2)
	fmt.Println("Final contador:", <-c3+<-c4)
}

func incrementor(s string) chan int {
	salida := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			salida <- 1
			fmt.Println(s, i)
		}
		close(salida)
	}()
	return salida
}

func extractor(c chan int) chan int {
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
