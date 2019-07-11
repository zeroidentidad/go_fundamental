package main

import "fmt"

type vehiculos interface{}

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
	boeing747 := avion{}
	boeing757 := avion{}
	boeing767 := avion{}
	sanger := bote{}
	nautique := bote{}
	malibu := bote{}
	manejados := []vehiculos{prius, tacoma, bmw528, boeing747, boeing757, boeing767, sanger, nautique, malibu}

	for key, value := range manejados {
		fmt.Println(key, " - ", value)
	}
}
