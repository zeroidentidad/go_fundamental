package main

import "fmt"

func main() {

	var miSaludo map[string]string
	fmt.Println(miSaludo)
	fmt.Println(miSaludo == nil)
}

// añade estas lineas:
/*
	miSaludo["Tim"] = "Good morning."
	miSaludo["Jenny"] = "Bonjour."
*/
// y obtendrás esto:
// panic: asignación a la entrada en el mapa nulo
