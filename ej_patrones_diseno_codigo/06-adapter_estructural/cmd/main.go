package main

import (
	"fmt"
	"log"

	"../adapter"
)

func main() {
	var t string

	fmt.Println("'bicicleta'\n'automovil'\nIngresar nombre transporte:")
	_, err := fmt.Scanln(&t)
	if err != nil {
		log.Fatalf("no se pudo leer el medio de transporte: %v", err)
	}
	transporteAd := adapter.Factory(t)
	transporteAd.Mover()
}
