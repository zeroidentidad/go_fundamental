package main

import "fmt"

type Animal struct {
	Nombre string
}

func (a Animal) Caminar() {
	fmt.Println("Caminar")
}

func (a Animal) Saltar() {
	fmt.Println("Saltar")
}

func main() {
	animal := Animal{"gato"}

	animal.Caminar()
	animal.Saltar()
}
