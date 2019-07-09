package main

import "fmt"

func main() {

	miSaludo := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	fmt.Println(miSaludo)

	// delete(miSaludo, 2)

	if valor, existe := miSaludo[2]; existe {
		fmt.Println("Ese valor existe.")
		fmt.Println("valor: ", valor)
		fmt.Println("existe: ", existe)
	} else {
		fmt.Println("Ese valor no existe.")
		fmt.Println("valor: ", valor)
		fmt.Println("existe: ", existe)
	}

	fmt.Println(miSaludo)
}
