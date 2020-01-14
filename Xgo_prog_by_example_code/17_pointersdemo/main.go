package main

import "fmt"

func main() {

	var x int

	x = 10
	fmt.Println(x)  // imprime valor de x
	fmt.Println(&x) // imprime dirección de x

	// declarar puntero
	var num *int
	val := new(int)

	num = new(int)
	*num = x // establecer valor

	val = &x // establecer dirección

	fmt.Println("== pointer num ==")
	fmt.Println(*num) // imprime valor de x
	fmt.Println(num)  // imprime dirección de x

	fmt.Println("== pointer val ==")
	fmt.Println(*val) // imprime valor de x
	fmt.Println(val)  // imprime dirección de x

}
