package main

import (
	"fmt"
	"regexp"
)

const refString = `[{ \"email\": \"email@ejemplo.com\" \"telefono\": 914467890},
{ \"email\": \"otro@dominio.com\" \"telefono\": 914467890}]`

func main() {

	// Patrón simplificado
	emailRegexp := regexp.MustCompile("[a-zA-Z0-9]{1,}@[a-zA-Z0-9]{1,}\\.[a-z]{1,}")
	primero := emailRegexp.FindString(refString)
	fmt.Println("Primero: ")
	fmt.Println(primero)

	todo := emailRegexp.FindAllString(refString, -1)
	fmt.Println("Todo: ")
	for _, val := range todo {
		fmt.Println(val)
	}

}
