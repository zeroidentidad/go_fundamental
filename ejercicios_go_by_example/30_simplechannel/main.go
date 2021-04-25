package main

import (
	"fmt"
)

func main() {

	fmt.Println("simple channel")

	// definir channel
	c := make(chan int)

	// run function en background
	go func() {
		fmt.Println("goroutine process")

		c <- 10 // escribir datos en un channel
	}()

	val := <-c // leer datos de un channel
	fmt.Printf("value: %d\n", val)
}
