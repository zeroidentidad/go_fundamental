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

	if val, existe := miSaludo[7]; existe {
		delete(miSaludo, 7)
		fmt.Println("valor: ", val)
		fmt.Println("existe: ", existe)
	} else {
		fmt.Println("Ese valor no existe.")
		fmt.Println("valor: ", val)
		fmt.Println("existe: ", existe)
	}

	fmt.Println(miSaludo)
}
