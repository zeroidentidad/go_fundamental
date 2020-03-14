package main

import (
	"fmt"
)

const x = 42
const y float64 = 43.2

type hotdog int
type hotcat float64

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	fmt.Printf("%T\n", hotdog(x))
	fmt.Printf("%T\n", hotdog(y))

	fmt.Printf("%T\n", hotcat(x))
	fmt.Printf("%T\n", hotcat(y))

}
