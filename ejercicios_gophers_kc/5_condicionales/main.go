package main

import (
	"fmt"
)

func main() {
	i := 10
	// Como definir un if - condicional
	if i == 10 {
		fmt.Println("I es 10")
	}
	j := 11
	// sugar syntax -- Atentos!!! -> Res coge el contexto del if
	// res no estaria disponible fuera del if unicamente dentro de el.
	if res := i + j; res > 20 {
		fmt.Println("Aprovado")
		fmt.Println(res)
	}

	res := i + j
	if res > 20 {
		fmt.Println("Aprovador")
		fmt.Println(res)
	} else {
		fmt.Println("Suspendido")
	}
	fmt.Println(res)

	if res > 20 {
		nombre := "Adrian"
		fmt.Println(nombre)
	}

	// fmt.Println(nombre) --> Error nombre no definido
	// Cuando estamos trabajando dentro de un if y creamos una variable dentro,
	// no estara disponible fuera del if.

}
