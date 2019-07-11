package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2) // agregar 2 rutinas al grupo de espera
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
	wg.Done()
}
