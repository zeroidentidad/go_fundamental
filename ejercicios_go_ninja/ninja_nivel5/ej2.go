package main

import (
	"fmt"
)

type persona struct {
	nombre     string
	apellido   string
	saboresFav []string
}

func main() {
	p1 := persona{
		nombre:   "Zero",
		apellido: "Ferrer",
		saboresFav: []string{
			"chocolate",
			"mantecado",
			"torta suiza",
		},
	}
	p2 := persona{
		nombre:   "Condor",
		apellido: "Pérez",
		saboresFav: []string{
			"fresa",
			"vainilla",
			"limón",
		},
	}

	m := map[string]persona{
		p1.apellido: p1,
		p2.apellido: p2,
	}

	for _, v := range m {
		fmt.Println(v.nombre)
		fmt.Println(v.apellido)
		for i, v := range v.saboresFav {
			fmt.Println(" ", i, v)
		}
		fmt.Println("-------")
	}
}
