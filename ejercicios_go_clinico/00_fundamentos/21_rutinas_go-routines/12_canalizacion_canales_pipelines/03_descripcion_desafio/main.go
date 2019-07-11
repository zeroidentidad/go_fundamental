package main

import (
	"fmt"
)

func main() {
	c := factorial(4)
	for n := range c {
		fmt.Println(n)
	}
}

func factorial(n int) chan int {
	salida := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		salida <- total
		close(salida)
	}()
	return salida
}

/*
DESAFIO # 1:
- Cambie el código de arriba para ejecutar 100 cálculos factoriales al mismo tiempo y en paralelo.
- Utilice el patrón de "pipeline" para lograr esto
*/
