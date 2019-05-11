package main

import "fmt"

func main() {
	saludarPersona("Jesus", 26)
}

func saludarPersona(nombre string, edad uint8) {
	//fmt.Println("Hola", nombre, "tienes", edad)
	fmt.Printf("Hola %s tienes %d años\n", nombre, edad)
}
