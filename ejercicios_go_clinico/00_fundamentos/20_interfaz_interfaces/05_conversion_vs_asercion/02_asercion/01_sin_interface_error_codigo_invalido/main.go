package main

import "fmt"

func main() {
	name := "Sydney"
	str, ok := name.(string)
	if ok {
		fmt.Printf("%q\n", str)
	} else {
		fmt.Printf("el valor no es una cadena\n")
	}
}
