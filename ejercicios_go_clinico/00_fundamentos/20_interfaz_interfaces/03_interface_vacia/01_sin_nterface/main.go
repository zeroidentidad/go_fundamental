package main

import "fmt"

type vehiculo struct {
	Asientos        int
	VelocidadMaxima int
	Color           string
}

type carro struct {
	vehiculo
	ruedas  int
	Puertas int
}

type avion struct {
	vehiculo
	Jet bool
}

type bote struct {
	vehiculo
	Longitud int
}

func main() {
	prius := carro{}
	tacoma := carro{}
	bmw528 := carro{}
	carros := []carro{prius, tacoma, bmw528}

	boeing747 := avion{}
	boeing757 := avion{}
	boeing767 := avion{}
	aviones := []avion{boeing747, boeing757, boeing767}

	sanger := bote{}
	nautique := bote{}
	malibu := bote{}
	botes := []bote{sanger, nautique, malibu}

	for key, value := range carros {
		fmt.Println(key, " - ", value)
	}

	for key, value := range aviones {
		fmt.Println(key, " - ", value)
	}

	for key, value := range botes {
		fmt.Println(key, " - ", value)
	}
}
