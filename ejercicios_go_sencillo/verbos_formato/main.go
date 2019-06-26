package main

import (
	"fmt"
)

func main() {
	var ancho, largo, area float64
	ancho = 7.5
	largo = 8.3
	area = ancho * largo
	fmt.Printf("cajas necesarias: %.2f \n", area/10.0)

	ancho = 6.5
	largo = 7.7
	area = ancho * largo
	fmt.Printf("cajas necesarias: %.2f \n", area/10.0)

}
