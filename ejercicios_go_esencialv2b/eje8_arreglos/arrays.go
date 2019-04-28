package main

import "fmt"

func main() {

	// forma definida
	var arreglo [2]string
	arreglo[0] = "Fruta"
	arreglo[1] = "Verduras"

	// otra forma o sintaxis
	arreglo1 := [2]string{"Fruta1", "Verduras1"}

	arreglo1[1] = "Tamalito"

	fmt.Println(arreglo, arreglo1)
}
