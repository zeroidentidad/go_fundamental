package main

import (
	"fmt"
)

func main() {
	a := 2

	pA := &a        // & Cogemos la dirección de memoria de a y la guardamos en pA
	fmt.Println(pA) // Imprimimos pA --> Dirección de memoria de a
	fmt.Println(&a)
	fmt.Println(pA == &a)
	fmt.Println(*pA) // * Accede dentro de pA (por lo que nos dara el valor de a)

	main2(pA)
	fmt.Println(a)
	pNil := returnNil()
	// Siempre que recivimos un puntero de una funcion realizamos un
	//if para ver si lo que contiene es nil
	// Cuidado con 3 party libraries
	if pNil == nil {
		fmt.Println("Esto es nil no podemos ejecutar por que sino nos saldra el panic")
	} else {
		fmt.Println(*pNil)
	}

}

// Si pasamos un puntero podemos modificar el contenido del puntero.
// *puntero **puntero 10<-0x...<-0x...
func main2(pA *int) {
	*pA = 10
}

func returnNil() *int {
	return nil
}
