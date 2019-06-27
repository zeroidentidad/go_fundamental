package main

import (
	"fmt"
)

func main() {

	var miInt int
	var miIntPointer *int
	miIntPointer = &miInt
	fmt.Println(miIntPointer)

	var miFloat float64
	var miFloatPointer *float64
	miFloatPointer = &miFloat
	fmt.Println(miFloatPointer)

	var miBool bool
	miBoolPointer := &miBool
	fmt.Println(miBoolPointer)
}
