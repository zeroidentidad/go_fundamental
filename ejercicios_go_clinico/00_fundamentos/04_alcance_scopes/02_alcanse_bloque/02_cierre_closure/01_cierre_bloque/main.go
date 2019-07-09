package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "El premio pertenece al que permanece en el ring."
		fmt.Println(y)
	}
	/* fmt.Println(y) */ // aqui fuera del alcance de y
}
