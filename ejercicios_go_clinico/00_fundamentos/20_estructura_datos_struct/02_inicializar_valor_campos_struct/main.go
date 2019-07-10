package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	edad     int
	ego      bool
}

func main() {
	p1 := persona{"James", "Bond", 20, true}
	p2 := persona{"Miss", "Moneypenny", 18, true}
	fmt.Println(p1.nombre, p1.apellido, p1.edad, p1.ego)
	fmt.Println(p2.nombre, p2.apellido, p2.edad, p2.ego)
}
