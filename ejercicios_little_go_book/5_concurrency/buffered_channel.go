package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int, 100)

	for {
		c <- rand.Int()
		fmt.Println(len(c))
		time.Sleep(time.Millisecond * 50)
	}
}
