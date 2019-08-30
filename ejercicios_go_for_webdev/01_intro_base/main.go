package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
}

type agenteSecreto struct {
	persona       // <-
	licenseToKill bool
}

func (p persona) hablar() {
	fmt.Println(p.nombre, p.apellido, `dice, "Buen dia, James."`)
}

func (sa agenteSecreto) hablar() {
	fmt.Println(sa.nombre, sa.apellido, `dice, "Sacudido, no revuelto."`)
}

type humano interface {
	hablar()
}

func decirAlgo(h humano) {
	h.hablar()
}

func main() {
	p1 := persona{
		"Miss",
		"Moneypenny",
	}

	sa1 := agenteSecreto{
		persona{
			"James",
			"Bond",
		},
		true,
	}

	decirAlgo(p1)
	decirAlgo(sa1)
}
