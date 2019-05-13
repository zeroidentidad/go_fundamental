package main

import "fmt"

type Persona struct {
	nombre string
	edad   int8
}

func (p Persona) saludar() {
	fmt.Println("Hola soy persona")
}

func (p *Persona) asignarEdad(i int8) {
	if i >= 0 {
		p.edad = i
	} else {
		fmt.Println("Edad no valida")
	}
}

type Numero int

func (n Numero) mostrar() {
	fmt.Println("Soy numero")
}

func main() {
	var persona Persona
	persona.saludar()
	persona.asignarEdad(26)
	fmt.Println(persona)
	var numero Numero
	numero.mostrar()
}
