package main

import "fmt"

// || Or
// && And
// ! Not

func main() {
	if 5 > 10 && 1 > 7 {
		fmt.Printf("En bloque verdadero")
	}
	fmt.Println("Fin 1")

	//isValid := true

	if edad, nombre := 15, "Alvarito"; edad < 14 {
		fmt.Println(nombre, "es un niño")
	} else if edad < 18 {
		fmt.Println(nombre, "un adolescente")
	} else if edad < 30 {
		fmt.Println("Es un adulto")
	} else {
		fmt.Println("Es un adulto grandecito")
	}

	/*
		if isValid {
			fmt.Println("Esto está en el bloque de verdadero")
			fmt.Println(isValid)
			if 5 < 10 {
				fmt.Println("5 es menor a 10")
			} else {
				fmt.Println("5 no es menor a 10")
			}
		} else {
			fmt.Println("Esto está en el bloque falso")
			fmt.Println(isValid)
		}
	*/
	fmt.Println("Fin 2")

}
