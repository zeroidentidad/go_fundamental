package main

import "fmt"

func main() {

	// declaraciones:
	var frutas = []string{"Manzanas", "Uvas", "Peras", "Melones"}
	var frutas2 []string

	fmt.Println(frutas)
	fmt.Println(len(frutas))

	frutas2 = append(frutas, "Peras")

	fmt.Println(frutas2)
	fmt.Println(len(frutas2))

	frutas = append(frutas, "Fresas")

	fmt.Println(frutas)
	fmt.Println(len(frutas))

	fmt.Println("Mostrando rango de valores:", frutas[2:3])
	fmt.Println("Mostrando rango de valores:", frutas[0:2])

}
