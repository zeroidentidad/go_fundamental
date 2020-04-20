package main

import "fmt"

type Animal struct {
	Nombre string
}

func (a Animal) Caminar() {
	fmt.Println("Caminar original")
}

func (a Animal) Saltar() {
	fmt.Println("Saltar")
}

type AnimalModificado struct {
	Animal
}

func (am AnimalModificado) Caminar() {
	fmt.Println("Caminar modificado")
}

func main() {
	animalOriginal := Animal{"gato"}
	animalOriginal.Caminar()
	animalOriginal.Saltar()

	fmt.Println("")

	animalModificado := AnimalModificado{Animal{"leopardo"}}
	animalModificado.Caminar()
	animalModificado.Saltar()
}
