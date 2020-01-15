package main

import (
	"fmt"
	"time"
)

func main() {

	l, err := time.LoadLocation("America/Mexico_City")
	if err != nil {
		panic(err)
	}
	t := time.Date(2020, 11, 30, 11, 10, 20, 0, l)
	fmt.Printf("Fecha predeterminada es: %v\n", t)

	// Sumar 3 días
	r1 := t.Add(72 * time.Hour)
	fmt.Printf("Fecha predeterminada +3D es: %v\n", r1)

	// Restar 3 días
	r1 = t.Add(-72 * time.Hour)
	fmt.Printf("Fecha predeterminada -3D es: %v\n", r1)

	// API comoda para agregar días/meses/años
	r1 = t.AddDate(1, 3, 2)
	fmt.Printf("Fecha predeterminada +1YR +3MTH +2D es: %v\n", r1)
}
