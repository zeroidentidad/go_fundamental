package main

import (
	"fmt"
)

type persona struct {
	Nombre   string
	Apellido string
	Edad     int
}

type doubleZero struct {
	persona
	Nombre            string
	LicenciaParaMatar bool
}

func main() {
	p1 := doubleZero{
		persona: persona{
			Nombre:   "James",
			Apellido: "Bond",
			Edad:     20,
		},
		Nombre:            "Double Zero Seven",
		LicenciaParaMatar: true,
	}

	p2 := doubleZero{
		persona: persona{
			Nombre:   "Miss",
			Apellido: "MoneyPenny",
			Edad:     19,
		},
		Nombre:            "If looks could kill",
		LicenciaParaMatar: false,
	}

	// los campos y métodos del tipo interno son promovidos al tipo externo
	fmt.Println(p1.Nombre, p1.persona.Nombre)
	fmt.Println(p2.Nombre, p2.persona.Nombre)
}
