package main

import (
	"fmt"
	"strconv"
)

func main() {

	c := incrementor(2)

	var contador int
	for n := range c {
		contador++
		fmt.Println(n)
	}

	fmt.Println("Contador final:", contador)
}

func incrementor(n int) chan string {
	c := make(chan string)
	terminado := make(chan bool)

	for i := 0; i < n; i++ {
		go func(i int) {
			for k := 0; k < 20; k++ {
				c <- fmt.Sprint("Proceso: "+strconv.Itoa(i)+" imprimiendo:", k)
			}
			terminado <- true
		}(i)
	}

	go func() {
		for i := 0; i < n; i++ {
			<-terminado
		}
		close(c)
	}()

	return c
}
