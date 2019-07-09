package main

import "fmt"

func main() {

	var miSaludo = make(map[string]string)
	miSaludo["Tim"] = "Good morning."
	miSaludo["Jenny"] = "Bonjour."

	fmt.Println(miSaludo)
}
