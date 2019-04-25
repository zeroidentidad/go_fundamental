
// declarar slice vacio, crear loop que inserte 10 valores
// iterar el slice y mostrar cada valor
// declarar slice de 5 string y asignar con valores string
// tomar slice del primer y segundo indice
// desplegar la posicion y valor de cada elemento en el nuevo slice

package main

import "fmt"

func main(){

	var numeros []int

	for i := 1 ; i <= 10; i++ {
		numeros = append(numeros, i*10)
	}

	for _, numero := range numeros {
		fmt.Println(numero)
	}
	fmt.Println("---")

	cadenas := []string{"cadena1", "cadena2", "cadena3", "cadena4", "cadena5"}

	for k, cadena := range cadenas {
		fmt.Printf("Indice: %d Cadena: %s\n", k, cadena)
	}
	fmt.Println("---")

	slice := cadenas[1:3]

	for l, cadena := range slice{
		fmt.Printf("Indice: %d Cadena: %s\n", l, cadena)
	}


}