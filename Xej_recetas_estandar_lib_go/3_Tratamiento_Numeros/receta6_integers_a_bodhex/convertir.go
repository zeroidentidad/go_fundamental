package main

import (
	"fmt"
	"strconv"
)

const bin = "10111"
const hex = "1A"
const oct = "12"
const dec = "10"
const floatNum = 16.123557

func main() {

	// Convierte el valor binario en hexadecimal
	v, _ := ConvertInt(bin, 2, 16)
	fmt.Printf("Valor binario %s convertido a hex: %s\n", bin, v)

	// Convierte el valor hexadecimal en decimal.
	v, _ = ConvertInt(hex, 16, 10)
	fmt.Printf("Valor hexadecimal %s convertido a dec: %s\n", hex, v)

	// Convierte el valor octal en hexadecimal
	v, _ = ConvertInt(oct, 8, 16)
	fmt.Printf("Valor octal %s convertido a hex: %s\n", oct, v)

	// Convierte el valor decimal en octal
	v, _ = ConvertInt(dec, 10, 8)
	fmt.Printf("Valor decimal %s convertido a oct: %s\n", dec, v)

	//... y asi se puede hacer cualquier otra conversión.

}

// ConvertInt convierte el valor de cadena dado de base
// a la base (toBase) definida.
func ConvertInt(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}
