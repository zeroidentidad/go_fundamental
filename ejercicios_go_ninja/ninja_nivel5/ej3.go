package main

import (
	"fmt"
)

type vehiculo struct {
	puertas int
	color   string
}

type camion struct {
	vehiculo
	cuatroRuedas bool
}

type sedan struct {
	vehiculo
	lujoso bool
}

func main() {
	c := camion{
		vehiculo: vehiculo{
			puertas: 2,
			color:   "blanco",
		},
		cuatroRuedas: true,
	}

	s := sedan{
		vehiculo: vehiculo{
			puertas: 4,
			color:   "negro",
		},
		lujoso: false,
	}

	fmt.Println(c)
	fmt.Println(s)
	fmt.Println(c.puertas)
	fmt.Println(s.puertas)
}
