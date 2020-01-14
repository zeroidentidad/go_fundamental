package main

import (
	"fmt"
)

func main() {

	fmt.Println("access mymodule...")
	var c int
	c = Add(10, 6)
	fmt.Printf("add()=%d\n", c)
	c = Subtract(5, 8)
	fmt.Printf("subtract()=%d\n", c)
	c = Multiply(5, 3)
	fmt.Printf("multiply()=%d\n", c)
}
