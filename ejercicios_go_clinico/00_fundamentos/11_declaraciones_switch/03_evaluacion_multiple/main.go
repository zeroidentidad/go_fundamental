package main

import "fmt"

func main() {
	switch "M" {
	case "Tim", "Jenny":
		fmt.Println("Que onda Tim o Jenny")
	case "Marcus", "Medhi":
		fmt.Println("En ambos su nombre inica con M")
	case "Julian", "Sushant":
		fmt.Println("Que onda Julian / Sushant")
	}
}
