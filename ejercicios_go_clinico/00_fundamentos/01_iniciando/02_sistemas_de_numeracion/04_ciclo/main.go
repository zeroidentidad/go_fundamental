package main

import "fmt"

func main() {
	for i := 1000000; i < 1000100; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i) //-> ciclo de secuencia de numeros del formatos decimal, binario y hexadecimal
	}
}
