package main

import (
	"fmt"

	"./pkgdespedir"
	"./pkgsaludar"
)

func main() {

	nombre := "Jesus"

	pkgsaludar.Saludar(nombre)

	pkgsaludar.MeVes = "Esto es un texto asignado desde el main"
	fmt.Println(pkgsaludar.MeVes)

	pkgdespedir.Despedirse(nombre)
}
