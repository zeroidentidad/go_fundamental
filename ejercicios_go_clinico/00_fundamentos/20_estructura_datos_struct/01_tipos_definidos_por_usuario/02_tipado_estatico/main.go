package main

import "fmt"

type foo int

func main() {
	var miEdad foo
	miEdad = 26
	fmt.Printf("%T %v \n", miEdad, miEdad)

	var tuEdad int
	tuEdad = 25
	fmt.Printf("%T %v \n", tuEdad, tuEdad)

	// esto no funciona:
	//	 fmt.Println(miEdad + tuEdad)

	// esta conversión funciona:
	//	 fmt.Println(int(miEdad) + tuEdad)
}
