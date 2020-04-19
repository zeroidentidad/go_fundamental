package main

import "fmt"

type Figura interface {
	Dibujar()
}

type Cuadrado struct{}

func (c *Cuadrado) Dibujar() {
	fmt.Println("Dibujando un Cuadrado.")
}

type Triangulo struct{}

func (t *Triangulo) Dibujar() {
	fmt.Println("Dibujando un Triángulo.")
}

func dibujarFigura(f Figura) {
	f.Dibujar()
}

func main() {
	cuadrado := Cuadrado{}
	dibujarFigura(&cuadrado)

	triangulo := Triangulo{}
	dibujarFigura(&triangulo)
}
