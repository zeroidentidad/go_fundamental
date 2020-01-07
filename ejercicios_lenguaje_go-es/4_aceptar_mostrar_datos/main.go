package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero1, numero2, resultado int
var leyenda string

func main() {
	fmt.Println("Ingrese numero 1: ")
	fmt.Scanln(&numero1)

	fmt.Println("Ingrese numero 2: ")
	fmt.Scanln(&numero2)

	fmt.Println("Descripcion: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	resultado = numero1 + numero2
	fmt.Println(leyenda, resultado)

}
