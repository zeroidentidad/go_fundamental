package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	terminado := make(chan bool)

	go func() {
		for i := 0; i < 100000; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		for n := range c {
			fmt.Println(n)
		}
		terminado <- true
	}()

	go func() {
		for n := range c {
			fmt.Println(n)
		}
		terminado <- true
	}()

	<-terminado
	<-terminado

}
