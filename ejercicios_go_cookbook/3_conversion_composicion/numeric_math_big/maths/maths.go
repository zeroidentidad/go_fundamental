package maths

import (
	"fmt"
	"math"
)

// algunas de las funciones en el paquete math
func Examples() {
	//sqrt Examples
	i := 25

	// i es un int, entonces se convierte
	result := math.Sqrt(float64(i))

	// sqrt de 25 == 5
	fmt.Println(result)

	// ceil redondea hacia arriba
	result = math.Ceil(9.5)
	fmt.Println(result)

	// floor redondea hacia bajo
	result = math.Floor(9.5)
	fmt.Println(result)

	// math almacena algunas constantes:
	fmt.Println("Pi:", math.Pi, "E:", math.E)
}
