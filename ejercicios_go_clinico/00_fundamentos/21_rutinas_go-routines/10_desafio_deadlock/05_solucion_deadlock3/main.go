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
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}

// recuerda cerrar tu canal
// Si no cierra su canal, recibirá este error.
// fatal error: all goroutines are asleep - deadlock!

// ************** IMPORTANTE **************
// NECESITAS VERSIÓN 1.5.2 O MAYOR
// de lo contrario recibirá este error
// fatal error: all goroutines are asleep - deadlock!
