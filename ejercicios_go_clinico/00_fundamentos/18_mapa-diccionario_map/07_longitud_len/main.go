package main

import "fmt"

func main() {

	miSaludo := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}

	miSaludo["Harleen"] = "Howdy"

	fmt.Println(len(miSaludo))
}
