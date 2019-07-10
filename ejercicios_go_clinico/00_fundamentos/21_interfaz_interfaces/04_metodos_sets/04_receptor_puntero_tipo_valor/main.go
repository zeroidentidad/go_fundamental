package main

import (
	"fmt"
	"math"
)

type circulo struct {
	radio float64
}

type forma interface {
	area() float64
}

func (c *circulo) area() float64 {
	return math.Pi * c.radio * c.radio
}

func info(s forma) {
	fmt.Println("area:", s.area())
}

func main() {
	c := circulo{5}
	info(c)
	// se necesita recibir puntero para el tipo implementado
}
