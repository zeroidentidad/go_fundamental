package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre   string
	Apellido string
	Edad     int
}

func main() {
	var p1 persona
	fmt.Println(p1.Nombre)
	fmt.Println(p1.Apellido)
	fmt.Println(p1.Edad)

	bs := []byte(`{"Nombre":"James", "Apellido":"Bond", "Edad":20}`)
	json.Unmarshal(bs, &p1)

	fmt.Println("--------------")
	fmt.Println(p1.Nombre)
	fmt.Println(p1.Apellido)
	fmt.Println(p1.Edad)
	fmt.Printf("%T \n", p1)
}
