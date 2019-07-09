package main

import "fmt"

func main() {
	saludo("Joane", "Doe")
}

func saludo(nombre string, apellido string) {
	fmt.Println(nombre, apellido)
}
