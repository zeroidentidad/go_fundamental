package main

import "fmt"

type Persona struct {
	Nombre   string
	Apellido string
	Edad     int
}

// Struct in Func
func (p *Persona) Saludar() {
	fmt.Printf("Nombre: %s, Apellido: %s, Edad: %d\n", p.Nombre, p.Apellido, p.Edad)
}

func main() {
	persona := Persona{"Jose", "Sánchez", 33}
	persona.Saludar()
}
