package main

import (
	"fmt"
)

func main() {

	var name string
	fmt.Println("Cual es tu nombre?")
	fmt.Scanf("%s\n", &name)

	var age int
	fmt.Println("Cual es tu edad?")
	fmt.Scanf("%d\n", &age)

	fmt.Printf("Hola %s de %d, fuiste hakiao xD\n", name, age)

}
