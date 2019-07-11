package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var contador int

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Contador final:", contador)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		x := contador
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		contador = x
		fmt.Println(s, i, "contador:", contador)
	}
	wg.Done()
}

// go run -race main.go
// vs
// go run main.go
