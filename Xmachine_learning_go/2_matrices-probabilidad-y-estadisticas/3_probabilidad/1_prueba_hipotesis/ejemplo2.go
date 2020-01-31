package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat"
)

func main() {

	// Definir valores observados y esperados. La mayoría de las veces,
	// estos vendrán de sus datos (visitas al sitio web, etc.).
	observed := []float64{48, 52}
	expected := []float64{50, 50}

	// Calcular la estadística de prueba de Chi cuadrado.
	chiSquare := stat.ChiSquare(observed, expected)

	fmt.Println(chiSquare)
}
