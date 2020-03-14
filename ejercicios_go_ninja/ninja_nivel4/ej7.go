package main

import (
	"fmt"
)

func main() {
	xs1 := []string{"Batman", "Jefe", "Vestido de negro"}
	xs2 := []string{"Robin", "Ayudante", "Vestido de colores"}
	fmt.Println(xs1)
	fmt.Println(xs2)

	xxs := [][]string{xs1, xs2}
	fmt.Println(xxs)

	for i, reg := range xxs {
		fmt.Println("Registro: ", i)
		for j, val := range reg {
			fmt.Printf("\tíndice de posición: %v\tvalor: \t %v \n", j, val)
		}
	}
}
