package main

import (
	"fmt"
	"time"
)

func main() {
	for n := range gen() {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	time.Sleep(1 * time.Minute)
}

// gen es un generador roto que filtrará una goroutine.
func gen() <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			ch <- n
			n++
		}
	}()
	return ch
}

//Context hace posible administrar una cadena de llamadas dentro de la misma ruta de llamada al señalar Done channel. del contexto.

//fuente: https://rakyll.org/leakingctx/
