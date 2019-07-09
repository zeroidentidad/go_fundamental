package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "El premio pertenece a quien quede en el ring."
		fmt.Println(y)
	}
	// fmt.Println(y) // fuera del alcance de y
}
