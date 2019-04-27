package main

import "fmt"

func main() {
	var micafe = Cafe{
		nombre: "normal",
		precio: 10,
		azucar: true,
		leche:  0,
	}

	var micafe2 = Cafe{
		"capuchino",
		15,
		true,
		50,
	}

	fmt.Println(micafe)
	fmt.Println(micafe2)
}

/* dato tipo cafe */
type Cafe struct {
	nombre string
	precio float64
	azucar bool
	leche  int
}
