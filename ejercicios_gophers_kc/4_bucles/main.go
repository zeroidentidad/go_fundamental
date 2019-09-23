package main

import (
	"fmt"
)

func main() {
	resultado := 0
	for i := 0; i < 1000; i++ {
		resultado++
	}
	fmt.Println(resultado)
	resultado = 1
	for resultado < 100 {
		fmt.Println(resultado)
		resultado += resultado
	}
	fmt.Println(resultado)
	for {

		return
	}

}
