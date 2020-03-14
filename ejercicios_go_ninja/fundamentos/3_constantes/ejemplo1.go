package main

import (
	"fmt"
)

const a = 42
const b = 42.32
const c = "Zero"

type nombre string

var d nombre

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	d = c
	fmt.Println(d)

}
