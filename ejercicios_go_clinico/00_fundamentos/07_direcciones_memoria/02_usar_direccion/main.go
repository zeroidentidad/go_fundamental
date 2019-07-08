package main

import "fmt"

const metrosAYardas float64 = 1.09361

func main() {
	var metros float64
	fmt.Print("Ingrese metros que nadó: ")
	fmt.Scan(&metros) //usar la variable en su direccion de memoria
	yardas := metros * metrosAYardas
	fmt.Println(metros, " metros son ", yardas, " yardas.")
}
