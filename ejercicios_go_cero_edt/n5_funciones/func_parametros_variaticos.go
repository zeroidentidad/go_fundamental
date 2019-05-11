package main

import "fmt"

func main() {
	saludarVarios(50, "Jesus", "Pedro", "Daniel", "Alvaro", "Diana")
}

func saludarVarios(edad uint8, nombres ...string) { // se permite uno variatico al final
	//fmt.Printf("%T\n", nombres)
	for _, v := range nombres {
		fmt.Println("Hola", v, "tiene", edad, "años")
	}
}
