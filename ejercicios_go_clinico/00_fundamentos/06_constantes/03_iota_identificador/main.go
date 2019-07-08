package main

import "fmt"

const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
)

// iota: identificador que se usa en las declaraciones const para simplificar las definiciones de números incrementales

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
