package main

import "fmt"

func main() {
	fmt.Println(saludo("Joane ", "Doe "))
}

func saludo(nombre, apellido string) (string, string) {
	return fmt.Sprint(nombre, apellido), fmt.Sprint(apellido, nombre)
}
