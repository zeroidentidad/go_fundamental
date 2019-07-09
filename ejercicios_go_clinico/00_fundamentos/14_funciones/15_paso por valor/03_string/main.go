package main

import "fmt"

func main() {

	nombre := "Jesus"
	fmt.Println(nombre) // Jesus

	cambiaMe(nombre)

	fmt.Println(nombre) // Jesus
}

func cambiaMe(z string) {
	fmt.Println(z) // Jesus
	z = "Rocky prro"
	fmt.Println(z) // Rocky prro
}
