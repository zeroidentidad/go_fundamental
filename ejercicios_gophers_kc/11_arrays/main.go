package main

import (
	"fmt"
)

func main() {
	// Arrays con numero de variables definido
	var hola [2]string
	hola[0] = "hola"
	hola[1] = "mundo"
	// hola[2] = "go" Petada --> Array unicamente de dos elementos
	fmt.Println(hola)
	var numeros []int
	numeros = append(numeros, 1)
	fmt.Println(numeros)
	letras := []string{"a", "b", "c"}
	fmt.Println(len(letras))
	fmt.Println(cap(letras))
	letras = append(letras, "d", "e")
	fmt.Println(len(letras))
	fmt.Println(cap(letras))
	letras = append(letras, "f", "g")
	fmt.Println(len(letras))
	fmt.Println(cap(letras))
	fmt.Println(letras)
	for _, letra := range letras {
		fmt.Println(letra)
	}

}
