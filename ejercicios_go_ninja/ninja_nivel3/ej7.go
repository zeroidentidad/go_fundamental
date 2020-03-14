package main

import (
	"fmt"
)

func main() {
	x := "Rsdfg"
	if x == "Robin" {
		fmt.Println(x)
	} else if x == "Batman" {
		fmt.Println("batbatbat", x)
	} else {
		fmt.Println("Ningún súper héroe")
	}
}
