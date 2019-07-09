package main

import "fmt"

func main() {

	saludos := []string{
		"Good morning!",
		"Bonjour!",
		"Buenos dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}

	for i, entradaActual := range saludos {
		fmt.Println(i, entradaActual)
	}

	for j := 0; j < len(saludos); j++ {
		fmt.Println(saludos[j])
	}

}
