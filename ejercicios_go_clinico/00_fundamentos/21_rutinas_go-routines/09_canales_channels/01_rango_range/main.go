package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i // llenar canal c en cada iterancion con valor de i
		}
		close(c) // cerrar llenado canal
	}()

	for n := range c { // recorrer canal y asignar valor a la variable receptora n
		fmt.Println(n)
	}
}
