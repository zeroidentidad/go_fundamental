package main

import (
	"fmt"
	"math"
)

type cuadrado struct {
	lado float64
}

// otra forma
type circulo struct {
	radio float64
}

type forma interface {
	area() float64
}

func (s cuadrado) area() float64 {
	return s.lado * s.lado
}

func (c circulo) area() float64 {
	return math.Pi * c.radio * c.radio
}

// que implementa la interface forma
func info(z forma) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := cuadrado{10}
	c := circulo{5}
	info(s)
	info(c)
}
