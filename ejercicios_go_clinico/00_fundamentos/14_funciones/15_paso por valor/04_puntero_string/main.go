package main

import "fmt"

func main() {

	nombre := "Jesus"
	fmt.Println(&nombre) // 0xc0000101e0

	cambiaMe(&nombre)

	fmt.Println(&nombre) //0xc0000101e0
	fmt.Println(nombre)  //Rocky
}

func cambiaMe(z *string) {
	fmt.Println(z)  // 0xc0000101e0
	fmt.Println(*z) // Jesus
	*z = "Rocky prro"
	fmt.Println(z)  // 0xc0000101e0
	fmt.Println(*z) // Rocky
}
