package main

import "fmt"

func main(){

	alumnos := 1

	fmt.Println("alumnos:")

	for alumnos <= 3 {
		fmt.Println(alumnos)

		alumnos = alumnos + 1 
	}

	fmt.Println("calificaciones:")

	for calificaciones := 7 ; calificaciones <= 9 ; calificaciones++ {
		fmt.Println(calificaciones)
	}
}