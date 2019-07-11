package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var contador int
var mutex sync.Mutex

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Contador final:", contador)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		mutex.Lock()
		contador++
		fmt.Println(s, i, "contador:", contador)
		mutex.Unlock()
	}
	wg.Done()
}

// go run -race main.go
// vs
// go run main.go
