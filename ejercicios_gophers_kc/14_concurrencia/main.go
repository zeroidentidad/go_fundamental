package main

import (
	"fmt"
	"time"
)

func run(num int, callback func(int)) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Gorutina %v %v\n", num, i)
	}
	callback(num)
}

func main() {
	timeStart := time.Now()
	numeroThreads := 1000
	callback := func(numeroGoRutina int) {
		numeroThreads--
		fmt.Printf("Ha finalizado la rutina: %v\n ", numeroGoRutina)
	}
	for i := 0; i < 1000; i++ {
		go run(i, callback)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	for {
		if numeroThreads == 0 {
			break
		}
	}
	fmt.Println(time.Now().Sub(timeStart))
}
