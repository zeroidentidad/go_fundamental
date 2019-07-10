package main

import "fmt"

type cuadrado struct {
	lado float64
}

func (z cuadrado) area() float64 {
	return z.lado * z.lado
}

func main() {
	s := cuadrado{10}
	fmt.Println("Area: ", s.area())
}
