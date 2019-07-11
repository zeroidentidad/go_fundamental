package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	age      int
}

func (p persona) nombreCompleto() string {
	return p.nombre + p.apellido
}

func main() {
	p1 := persona{"James", "Bond", 20}
	p2 := persona{"Miss", "Moneypenny", 18}
	fmt.Println(p1.nombreCompleto())
	fmt.Println(p2.nombreCompleto())
}
