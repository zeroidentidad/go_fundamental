package main

import "fmt"

type cuadrado struct {
	lado float64
}

// recibe tipo z, nombre func area, retorno float
func (z cuadrado) area() float64 {
	return z.lado * z.lado
}

type forma interface {
	area() float64
}

// func haciendo uso de la interfaz como tipo forma y propiedad func area()
func info(z forma) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	//slice tipo cuadrado, propiedad lado
	s := cuadrado{10}
	info(s)
}
