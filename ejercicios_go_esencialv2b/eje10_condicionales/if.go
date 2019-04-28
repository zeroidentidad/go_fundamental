package main

import "fmt"

func main() {
	puntos := 500

	if puntos < 100 {
		fmt.Println("puntos incorrectos")
	} else if puntos == 100 {
		fmt.Println("puntos correctos :)")
	} else {
		fmt.Println("tus puntos son:", puntos)
	}
}
