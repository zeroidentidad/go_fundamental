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

	fmt.Println(<-c)
}

// ¿Por qué esto solo imprime cero?
// ¿Y qué puedes hacer para imprimir todos los números del 0 al 9?
