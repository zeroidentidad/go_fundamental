package main

import (
	"curso-go/9_estructuras/model"
	"fmt"
)

func main() {
	motor := model.Motor{
		Cilindrada:   2000,
		NumCilindros: 1,
	}
	coche := model.Coche{
		Motor:        nil,
		NumeroRuedas: 6,
	}
	fmt.Println(motor.NumCilindros)
	// Trabajando con punteros nunca poner el asterisco ya que el punto accede
	// dentro de la refencia a memoria.
	if coche.Motor != nil {
		fmt.Println(coche.Motor.Cilindrada)
	} else {
		fmt.Println("Coche de los picapiedra")
	}
}
