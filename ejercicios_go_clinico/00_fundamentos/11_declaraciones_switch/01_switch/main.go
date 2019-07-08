package main

import "fmt"

func main() {
	switch "Mao" {
	case "Daniel":
		fmt.Println("Que onda Daniel")
	case "Medhi":
		fmt.Println("Que onda Medhi")
	case "Jenny":
		fmt.Println("Que onda Jenny")
	default:
		fmt.Println("No tienes amigos? :v")
	}
}

/*
  sin fallthrough por defecto : caso de continuar desde aqui
  fallthrough es opcional
  -- se puede especificar fallthrough indicandolo explicitamente
  -- break no es necesario como en otros lenguajes
*/
