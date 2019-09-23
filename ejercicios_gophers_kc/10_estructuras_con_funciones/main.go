package main

import (
	"curso-go/9_estructuras/model"
	"fmt"
)

func main() {

	coche := model.NewCoche(2, 600, 2)
	fmt.Println(coche)
	coche.IncrementarPotencia(100)
	coche.PincharRueda()
	coche.PincharRueda()
	coche.PincharRueda()
	fmt.Println(coche.NumeroRuedas)
	fmt.Println(coche.Arranca())

}
