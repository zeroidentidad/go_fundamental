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
	LicenciaParaMatar bool
}

func main() {
	p1 := doubleZero{
		persona: persona{
			Nombre:   "James",
			Apellido: "Bond",
			Edad:     20,
		},
		LicenciaParaMatar: true,
	}

	p2 := doubleZero{
		persona: persona{
			Nombre:   "Miss",
			Apellido: "MoneyPenny",
			Edad:     19,
		},
		LicenciaParaMatar: false,
	}

	fmt.Println(p1.Nombre, p1.Apellido, p1.Edad, p1.LicenciaParaMatar)
	fmt.Println(p2.Nombre, p2.Apellido, p2.Edad, p2.LicenciaParaMatar)
}
