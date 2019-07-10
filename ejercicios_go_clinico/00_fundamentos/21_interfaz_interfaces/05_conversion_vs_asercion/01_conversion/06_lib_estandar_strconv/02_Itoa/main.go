package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 12
	y := "Tengo tantos: " + strconv.Itoa(x) // equivalente a FormatInt()
	fmt.Println(y)
	//	fmt.Println("Tengo tantos: ", strconv.Itoa(x), x)
}
