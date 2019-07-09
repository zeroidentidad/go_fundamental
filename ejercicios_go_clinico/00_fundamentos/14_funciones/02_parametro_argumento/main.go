package main

import "fmt"

func main() {
	saludo("Joane")
	saludo("John")
}

func saludo(name string) {
	fmt.Println("Hola", name)
}

// saludo se declara con un parametro
// cuando se llama saludo, pasa un argumento
