package main

import "fmt"

func main() {
	fmt.Println(saludo("Joane ", "Doe"))
}

func saludo(nombre, apellidos string) string {
	return fmt.Sprint(nombre, apellidos)
}
