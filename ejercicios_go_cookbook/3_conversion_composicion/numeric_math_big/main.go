package main

import (
	"fmt"

	"./maths"
)

func main() {
	maths.Examples()

	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", maths.Fib(i))
	}
	fmt.Println()
}
