package main

import "fmt"

func main() {
	var name interface{} = "Sydney"
	str, ok := name.(string)
	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Printf("el valor no es una cadena\n")
	}
}
