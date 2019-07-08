package main

import "fmt"

func main() {
	switch "Marcus" {
	case "Tim":
		fmt.Println("Que onda Tim")
	case "Jenny":
		fmt.Println("Que onda Jenny")
	case "Marcus":
		fmt.Println("Que onda Marcus")
		fallthrough // caer aqui
	case "Medhi":
		fmt.Println("Que onda Medhi")
		fallthrough // a travez de aqui tambien
	case "Julian":
		fmt.Println("Que onda Julian")
	case "Sushant":
		fmt.Println("Que onda Sushant")
	}
}
