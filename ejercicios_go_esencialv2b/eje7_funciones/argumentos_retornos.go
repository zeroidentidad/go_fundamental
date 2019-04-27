package main

import "fmt"

func main() {
	fmt.Println(saludo())
	fmt.Println(saludo2())
}

func saludo() (string, string) {

	var a string
	var b string

	a = "Hola"
	b = "mundo"

	return a, b
}

func saludo2() (a string, b int) {

	a = "Hola"
	b = 26

	return
}
