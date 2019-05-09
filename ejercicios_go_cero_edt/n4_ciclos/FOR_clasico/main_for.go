package main

import "fmt"

func main() {
	for indice := 0; indice < 5; indice++ {
		fmt.Println(indice)
	}

	fmt.Println("--------------------")

	for indice := 5; indice >= 0; indice-- {
		if indice == 3 {
			fmt.Println("salto y continuar")
			continue
		}
		fmt.Println(indice)
	}

	fmt.Println("--------------------")

	for indice := 5; indice >= 0; indice-- {
		if indice == 2 {
			fmt.Println("interrupcion y terminar")
			break
		}
		fmt.Println(indice)
	}

	fmt.Println("--------------------")

	// var matriz [3][3]
	matriz := [3][3]int{}
	valor := 0

	for i := 0; i < len(matriz); i++ {
		for j := 0; j < 3; j++ {
			valor++
			matriz[i][j] = valor
		}
	}

	for fila := 0; fila < len(matriz); fila++ {
		for columna := 0; columna < 3; columna++ {
			fmt.Println("Fila-Columna:", matriz[fila][columna]) //, "Columna-Fila:", matriz[columna][fila]
		}
	}

}
