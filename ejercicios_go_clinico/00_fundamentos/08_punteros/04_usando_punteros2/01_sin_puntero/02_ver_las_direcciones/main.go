package main

import "fmt"

func zero(z int) {
	fmt.Printf("%p\n", &z) // direccion en func zero
	fmt.Println(&z)        // direccion en func zero
	z = 0
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x) // direccion en main
	fmt.Println(&x)        // direccion en main
	zero(x)
	fmt.Println(x) // x es aun 5
}
