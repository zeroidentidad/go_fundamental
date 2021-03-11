package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(randomNum())
}

func randomNum() int {
	//semilla a nivel nanosegundo
	rand.Seed(time.Now().UTC().UnixNano())

	return rand.Intn(1000) // valor entre 0 - 1000
}
