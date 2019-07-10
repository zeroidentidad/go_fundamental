package main

import (
	"fmt"
)

type persona struct {
	Nombre string
	Edad   int
}

type doubleZero struct {
	persona
	LicenciaParaMatar bool
}

func (p persona) Saludar() {
	fmt.Println("Solo soy una persona normal.")
}

func (dz doubleZero) Saludar() {
	fmt.Println("Miss Moneypenny, que bueno verte.")
}

func main() {
	p1 := persona{
		Nombre: "Ian Flemming",
		Edad:   44,
	}

	p2 := doubleZero{
		persona: persona{
			Nombre: "James Bond",
			Edad:   23,
		},
		LicenciaParaMatar: true,
	}
	p1.Saludar()
	p2.Saludar()
	p2.persona.Saludar()
}
