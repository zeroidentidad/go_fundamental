package main

import "fmt"

type Persona struct {
	Nombre   string
	Apellido string
	Edad     int
}

type Empleado struct {
	Persona
	Legajo string
}

func (p *Persona) DatosBase() {
	fmt.Printf("Nombre: %s, Apellido: %s, Edad: %d\n", p.Nombre, p.Apellido, p.Edad)
}

func (e *Empleado) DatosAdicionales() {
	fmt.Printf("El Empleado %s %s tiene el Legajo: %s\n", e.Nombre, e.Apellido, e.Legajo)
}

func main() {
	empleado := Empleado{
		Persona{"Jose", "Sánchez", 33},
		"A1234",
	}
	empleado.DatosBase()
	empleado.DatosAdicionales()
}
