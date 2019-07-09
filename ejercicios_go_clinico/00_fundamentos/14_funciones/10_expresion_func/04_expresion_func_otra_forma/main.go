package main

import "fmt"

func hacerSaludo() func() string {
	return func() string {
		return "Hello world!"
	}
}

func main() {
	saludo := hacerSaludo()
	fmt.Println(saludo())
}
