package main

import (
	"fmt"
)

func main() {
	var registros [][]string
	// estudiante 1
	estudiante1 := make([]string, 4)
	estudiante1[0] = "Foster"
	estudiante1[1] = "Nathan"
	estudiante1[2] = "100.00"
	estudiante1[3] = "74.00"
	// guardar el registro
	registros = append(registros, estudiante1)
	// estudiante 2
	estudiante2 := make([]string, 4)
	estudiante2[0] = "Gomez"
	estudiante2[1] = "Lisa"
	estudiante2[2] = "92.00"
	estudiante2[3] = "96.00"
	// guardar el registro
	registros = append(registros, estudiante2)
	// print
	fmt.Println(registros)
}
