package main

import "fmt"

type B struct{}

func (b B) Saludar() {
	fmt.Println("Saludando desde B")
}

type A struct {
	B
}

func main() {
	a := A{}

	a.Saludar()
}
