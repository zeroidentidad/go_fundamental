package main

import "fmt"

func main() {
	n := hashCubo("Go", 12)
	fmt.Println(n)
}

func hashCubo(palabra string, cubos int) int {
	letra := int(palabra[0])
	cubo := letra % cubos
	return cubo
}
