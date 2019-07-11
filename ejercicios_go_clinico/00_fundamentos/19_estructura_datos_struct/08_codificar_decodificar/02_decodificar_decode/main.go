package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type person struct {
	Nombre       string
	Apellido     string
	Edad         int
	notExportado int
}

func main() {
	var p1 person
	rdr := strings.NewReader(`{"Nombre":"James", "Apellido":"Bond", "Edad":20}`)
	json.NewDecoder(rdr).Decode(&p1) //salida sin formato json, solo valores string

	fmt.Println(p1.Nombre)
	fmt.Println(p1.Apellido)
	fmt.Println(p1.Edad)
}
