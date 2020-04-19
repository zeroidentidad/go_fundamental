package main

import (
	"fmt"
)

// tipo de dato especializado
type Int int

func main() {
	var num1 Int = 20
	var num2 Int = 25

	fmt.Printf("¿%v es mayor que %v?. Respuesta: %v\n", num1, num2, num1.EsMayorQue(num2))
}

func (n Int) EsMayorQue(n2 Int) bool {
	return n > n2
}
