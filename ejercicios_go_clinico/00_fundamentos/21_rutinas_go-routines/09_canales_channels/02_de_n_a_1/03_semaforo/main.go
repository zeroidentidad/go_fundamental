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

	go func() {
		<-terminado
		<-terminado
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
