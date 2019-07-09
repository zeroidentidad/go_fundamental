package main

import "fmt"

func main() {

	saludo := func() {
		fmt.Println("Hello world!")
	}

	saludo()
	fmt.Printf("%T\n", saludo)
}
