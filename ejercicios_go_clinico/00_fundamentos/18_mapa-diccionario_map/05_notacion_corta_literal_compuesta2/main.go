package main

import "fmt"

func main() {

	miSaludo := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}

	fmt.Println(miSaludo["Jenny"])
}
