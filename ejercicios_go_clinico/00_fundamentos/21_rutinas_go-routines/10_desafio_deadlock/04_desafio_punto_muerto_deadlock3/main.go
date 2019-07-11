package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	for {
		fmt.Println(<-c)
	}
}

// Esto imprime el número, pero hemos recibido este error:
// fatal error: all goroutines are asleep - deadlock!
// ¿Dónde está el punto muerto?
// ¿Por qué todos los goroutines están dormidos?
// ¿Como podemos arreglar esto?
