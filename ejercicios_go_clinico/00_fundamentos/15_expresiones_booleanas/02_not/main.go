package main

import "fmt"

func main() {

	if !true {
		fmt.Println("Esto no se ejecuta")
	}

	if !false {
		fmt.Println("Esto se ejecuta")
	}

}
