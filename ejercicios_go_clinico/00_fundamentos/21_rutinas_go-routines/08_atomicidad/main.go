package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var contador int64

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Contador final:", contador)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		atomic.AddInt64(&contador, 1)
		fmt.Println(s, i, "contador:", contador)
	}
	wg.Done()
}

// go run -race main.go
// vs
// go run main.go

//

/* "sync/atomic"
proporciona primitivas de memoria atómica de bajo nivel útiles para implementar algoritmos de sincronización
*/
