package main

import "fmt"

func main() {
	alumnos("jesus", "26", "tabasco")
}

func alumnos(alumno ...string) { // nombre string, edad int

	for _, valor := range alumno {
		fmt.Println(valor)
	}

}
