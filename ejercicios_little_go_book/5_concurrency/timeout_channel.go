package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	c := make(chan int)

	for {
		select {
		case c <- rand.Int():
		case t := <-time.After(time.Millisecond * 100):
			fmt.Println("timed out at", t)
		}
		time.Sleep(time.Millisecond * 50)
	}

	/*for {
		select {
		case data := <-c:
			fmt.Printf("worker %d got %d\n", w.id, data)
		case <-time.After(time.Millisecond * 10):
			fmt.Println("Break time")
			time.Sleep(time.Second)
		}
	}*/
}

func after(d time.Duration) chan bool { // like time.After works
	c := make(chan bool)
	go func() {
		time.Sleep(d)
		c <- true
	}()
	return c
}
