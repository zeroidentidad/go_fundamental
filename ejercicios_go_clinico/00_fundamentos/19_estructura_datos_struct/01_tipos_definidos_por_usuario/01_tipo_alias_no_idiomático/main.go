package main

import "fmt"

type foo int

func main() {
	var miEdad foo
	miEdad = 44
	fmt.Printf("%T %v \n", miEdad, miEdad)
}
