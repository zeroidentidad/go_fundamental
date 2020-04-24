package main

import "fmt"

type celsius float64
type kelvin float64

func main() {
	var k kelvin = 294.0
	c := kelvinToCelsius(k)

	fmt.Println(c)
	fmt.Println(k.celsius())
}

// funcion
func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

// metodo
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}
