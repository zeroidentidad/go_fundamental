package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	terminado := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		terminado <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		terminado <- true
	}()

	// bloqueamos aquí hasta terminado <- true
	<-terminado
	<-terminado
	close(c)

	// para desbloquear arriba
	// necesitamos quitar los valores de chan c aquí
	// pero nunca llegamos aquí, porque estamos bloqueados arriba
	for n := range c {
		fmt.Println(n)
	}
}
