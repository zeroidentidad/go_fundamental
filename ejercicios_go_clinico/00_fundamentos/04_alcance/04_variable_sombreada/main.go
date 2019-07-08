package main

import "fmt"

func max(x int) int {
	return 42 + x
}

func main() {
	max := max(7)
	fmt.Println(max) // max es ahora el resultado, no la función.
}

// no hacer esto mala práctica de codificación para sombrear variables
