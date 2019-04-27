package main

import (
	"fmt"
)

func main() {

	const estatica bool = true // forma estricta
	const estatica2 = false    // forma asumiendo tipo de dato

	/*
		estatica = false
	*/

	fmt.Println(estatica, estatica2)
}
