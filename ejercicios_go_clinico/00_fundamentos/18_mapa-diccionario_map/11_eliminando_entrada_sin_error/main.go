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
	delete(miSaludo, 7) // no genera conflicto el que no exista
	fmt.Println(miSaludo)
}
