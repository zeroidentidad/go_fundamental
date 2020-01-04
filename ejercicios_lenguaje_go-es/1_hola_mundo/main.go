package main

import (
	"fmt"
)

var archivo string
var suma, resta int
var multi, division float32

func main() {
	fmt.Println("Hola Mundo")
	suma, resta, multi, division, archivo = mostrarArchivo(3, 4)
	fmt.Printf("Suma %d", suma)
	fmt.Printf(" Resta %d", resta)
	fmt.Printf(" Multiplicacion %.2f", multi)
	fmt.Printf(" Division ", division)
	fmt.Printf("\n Archivo " + archivo)
}

func mostrarArchivo(numero1 int, numero2 int) (int, int, float32, float32, string) {
	return numero1 + numero2, numero1 - numero2, float32(numero1 * numero2), float32(numero1 / numero2), "Mi Archivo"
}
