package main

import "fmt"

func main() {

	tigres := 15

	for i := 0; i < tigres; i++ {

		if i == 1 {
			fmt.Println(i, "triste tigre")
		} else {
			fmt.Println(i, "tristes tigres")
		}

	}
}
