package main

import "fmt"

func hola() {
	fmt.Print("hola ")
}

func mundo() {
	fmt.Println("mundo prros")
}

func main() {
	defer mundo()
	hola()
}
