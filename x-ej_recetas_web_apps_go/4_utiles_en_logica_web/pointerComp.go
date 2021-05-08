package main

import (
	"fmt"
)

type persona struct {
	nombre    string
	edad      int
	domicilio string
}

func main() {
	var a int = 100
	var p *int = &a

	fmt.Println(&a)
	fmt.Println(p)
	fmt.Println(*p) // desreferenciar

	jesus := &persona{
		"Jesus",
		27,
		"Mexico",
	}

	(*jesus).domicilio = "Monterrey"

	fmt.Printf("%T\n", jesus)
	fmt.Printf("%v\n", jesus)

	superpoderes := &[4]string{"volar", "super fuerza", "invisibilidad", "control tiempo"}
	fmt.Println(superpoderes)

	rambo := persona{
		nombre: "Rambo",
		edad:   10,
	}

	rambo2 := &persona{
		nombre: "Rambo 2",
		edad:   10,
	}

	birthday(&rambo)
	fmt.Println(rambo)

	rambo.birthday()
	fmt.Println(rambo)

	rambo2.birthday()
	fmt.Println(rambo2)
}

func birthday(p *persona) {
	p.edad++
}

func (p *persona) birthday() {
	p.edad++
}
