package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		fmt.Println("processing 1")
	}()

	fmt.Println("start")
	time.Sleep(time.Millisecond * 5) // for demo porpuse
	go process()
	time.Sleep(time.Millisecond * 10) // for demo porpuse
	fmt.Println("done")
}

func process() {
	fmt.Println("processing 2")
	time.Sleep(time.Millisecond * 5) // for demo porpuse
}
