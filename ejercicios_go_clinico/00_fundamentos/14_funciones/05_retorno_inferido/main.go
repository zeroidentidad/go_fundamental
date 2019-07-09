package main

import "fmt"

func main() {
	fmt.Println(greet("Joane ", "Doe"))
}

func greet(nombre string, apellido string) (s string) {
	s = fmt.Sprint(nombre, apellido)
	return
}
