package main

import "fmt"

type Saludador interface {
	Saludar()
}

type B struct{}

func (b B) Saludar() {
	fmt.Println("Saludando desde B")
}

type A struct {
	Saludador
}

func main() {
	a := A{B{}}

	a.Saludar()
}
