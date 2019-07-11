package main

import (
	"fmt"
)

func main() {
	f := factorial(4)
	fmt.Println("Total:", f)
}

func factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

/*
DESAFIO # 1:
- Utilizar goroutines y canales para calcular factorial.

DESAFIO # 2:
- ¿Por qué podrías querer usar goroutines y canales para calcular factorial?
-- Formule su propia respuesta, luego publique esa respuesta en esta área de discusión: https://goo.gl/flGsyX
-- Lea algunas de las otras respuestas en el área de discusión para ver las razones de otros
*/
