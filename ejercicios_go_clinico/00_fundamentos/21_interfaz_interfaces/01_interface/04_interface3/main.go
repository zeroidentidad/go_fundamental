package main

import (
	"fmt"
	"math"
)

type circulo struct {
	radio float64
}

type cuadrado struct {
	lado float64
}

type forma interface {
	area() float64
}

func (c circulo) area() float64 {
	return math.Pi * c.radio * c.radio
}

func (s cuadrado) area() float64 {
	return s.lado * s.lado
}

func info(z forma) {
	fmt.Println(z)
	fmt.Println(z.area())
}

// un nuevo método que toma el TIPO DE INTERFAZ forma
func totalArea(formas ...forma) float64 {
	var area float64
	for _, s := range formas {
		area += s.area()
	}
	return area
}

func main() {
	s := cuadrado{10}
	c := circulo{5}
	info(s)
	info(c)
	fmt.Println("Area total: ", totalArea(c, s))
}
