package main

import (
	"fmt"
	"strconv"
)

const bin = "00001"
const hex = "2f"
const intString = "12"
const floatString = "12.3"

func main() {

	// Decimales
	res, err := strconv.Atoi(intString)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Resultado entero: %d\n", res)

	// Analizando hexadecimales
	res64, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Resultado hexadecimal: %d\n", res64)

	// Analizando valores binario
	resBin, err := strconv.ParseInt(bin, 2, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Resultado binario: %d\n", resBin)

	// Analizando puntos flotantes
	resFloat, err := strconv.ParseFloat(floatString, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Resultado flotante: %.5f\n", resFloat)

}
