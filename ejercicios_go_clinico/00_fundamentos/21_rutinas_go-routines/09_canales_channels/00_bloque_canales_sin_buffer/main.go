package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) // chan keyword

	go func() {
		for i := 0; i < 10; i++ {
			c <- i // llenar canal c en cada iterancion con valor de i
		}
	}()

	go func() {
		for {
			fmt.Println(<-c) // vaciar canal a la salida actual
		}
	}()

	time.Sleep(time.Second)
}
