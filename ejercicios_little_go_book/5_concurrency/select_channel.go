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
		// optional code here
		default:
			// this can be left empty to silently drop the data
			fmt.Println("dropped")
		}
		time.Sleep(time.Millisecond * 50)
	}
}
