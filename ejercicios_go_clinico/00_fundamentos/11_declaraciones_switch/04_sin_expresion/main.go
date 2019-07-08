package main

import "fmt"

func main() {

	nombreAmigo := "Ma"

	switch {
	case len(nombreAmigo) == 2:
		fmt.Println("Que onda mi amigo con nombre de longitud 2")
	case nombreAmigo == "Tim":
		fmt.Println("Que onda Tim")
	case nombreAmigo == "Jenny":
		fmt.Println("Que onda Jenny")
	case nombreAmigo == "Marcus", nombreAmigo == "Medhi":
		fmt.Println("Tu nombre es Marcus o Medhi")
	case nombreAmigo == "Julian":
		fmt.Println("Que onda Julian")
	case nombreAmigo == "Sushant":
		fmt.Println("Que onda Sushant")
	default:
		fmt.Println("nada conincide, este es el predeterminado")
	}
}

/*
expresión no es necesaria
   - Si no se proporciona una expresión, va a verificar el primer caso que sean verdadero
   - hace que el switch funcione como if/if else/else
   los casos pueden ser expresiones
*/
