package main

import (
	"fmt"

	"./simplemath"
)

func main() {

	fmt.Println("access mymodule...")
	var c int

	c = simplemath.Add(5, 8)
	fmt.Printf("add()=%d\n", c)

	c = simplemath.Subtract(5, 8)
	fmt.Printf("subtract()=%d\n", c)

	c = simplemath.Multiply(5, 3)
	fmt.Printf("multiply()=%d\n", c)
}
