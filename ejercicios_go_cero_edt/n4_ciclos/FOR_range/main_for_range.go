package main

import "fmt"

func main() {
	numeros := []string{"cero", "uno", "dos", "tres"} // slice

	nombres := map[string]string{"es": "España", "mex": "Mexico", "co": "Colombia"} // mapa forma corta

	for indice, valor := range numeros {
		fmt.Println(indice, valor)
	}

	fmt.Println("-----------------")

	for _, valor := range numeros {
		fmt.Println(valor)
	}

	fmt.Println("-----------------")

	for indice := range numeros {
		fmt.Println(indice)
	}

	fmt.Println("-----------------")

	for indice, valor := range nombres {
		fmt.Println(indice, valor)
	}

	fmt.Println("-----------------")

	frase := "Hola mundo"
	for posicion, letra := range frase {
		fmt.Println(posicion, string(letra), letra)
	}

	fmt.Println("-----------------")

	for _, letra := range frase {
		fmt.Println(string(letra))
	}

	fmt.Println("-----------------")

	for _, entero := range []int{12, 45, 23, 56} {
		fmt.Println(entero)
	}

}
