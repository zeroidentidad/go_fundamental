package main

import "fmt"

//  switch en types
// - normalmente cambiamos el valor de la variable
// - go le permite activar el tipo de variable

type contact struct {
	saludo string
	nombre string
}

// SwitchOnType trabaja con interfaces
// sobre las interfaces más adelante
func SwitchOnType(x interface{}) {
	switch x.(type) { // esto es una afirmación; afirmando, "x es de este tipo"
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")

	}
}

func main() {
	SwitchOnType(7)
	SwitchOnType("McLeod")
	var t = contact{"Que bueno verte,", "Tim"}
	SwitchOnType(t)
	SwitchOnType(t.saludo)
	SwitchOnType(t.nombre)
}
