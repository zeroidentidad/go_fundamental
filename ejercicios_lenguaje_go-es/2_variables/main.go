package main

import (
	"fmt"
	"strconv"
)

var numero int
var texto string
var status bool = true

func main() {
	numero2, numero3, numero4, texto, status := 2, 5, 67, "Este es mi texto", false

	var moneda float64 = 0

	numero2 = int(moneda)

	texto = strconv.Itoa(int(moneda))

	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(texto)
	fmt.Println(status)

	MostrarStatus()
}

func MostrarStatus() {
	fmt.Println(status)
}
