package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int64
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go incrementor("1")
	go incrementor("2")
	wg.Wait()
	fmt.Println("Contador final:", count)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		atomic.AddInt64(&count, 1)
		fmt.Println("Proceso: "+s+" imprimiendo:", i)
	}
	wg.Done()
}

/*
DESAFIO # 1
- Tome el código de arriba y cámbielo para usar canales en lugar de grupos de espera y atomicidad
*/
