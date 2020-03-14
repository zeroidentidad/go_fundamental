package main

import (
	"fmt"
)

//Creamos VALORES de ciertos TIPOS que se almacenan en VARIABLES
//las VARIABLES tienen IDENTIFICADORES

var x int

type persona struct {
	nombre   string
	apellido string
}

type foo int

var y foo

func main() {
	y = 42
	p1 := persona{
		nombre:   "Zero",
		apellido: "Ferrer",
	}
	fmt.Println(p1)
	fmt.Printf("%T", y)
}
