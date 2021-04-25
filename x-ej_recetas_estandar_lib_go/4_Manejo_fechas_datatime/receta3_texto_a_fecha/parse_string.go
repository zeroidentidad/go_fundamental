package main

import (
	"fmt"
	"time"
)

func main() {

	// Si la zona horaria no está definida, la función Parse
	// devuelve la hora en la zona horaria UTC.
	t, err := time.Parse("2/1/2006", "31/7/2015")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	// Si se da una zona horaria, entonces
	// se analiza en una zona horaria determinada
	t, err = time.Parse("2/1/2006  3:04 PM MST", "31/7/2015  1:25 AM DST")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	// Tener en cuenta que ParseInLocation
	// analiza el tiempo en una ubicación determinada, si la
	// la cadena no contiene la definición de zona horaria
	t, err = time.ParseInLocation("2/1/2006  3:04 PM ", "31/7/2015  1:25 AM ", time.Local)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

}
