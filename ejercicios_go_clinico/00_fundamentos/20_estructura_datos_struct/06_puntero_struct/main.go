package main

import "fmt"

type persona struct {
	nombre string
	edad   int
}

func main() {
	p1 := &persona{"James", 20}
	fmt.Println(p1)
	fmt.Printf("%T\n", p1)
	fmt.Println(p1.nombre)
	fmt.Println(p1.edad)
}
