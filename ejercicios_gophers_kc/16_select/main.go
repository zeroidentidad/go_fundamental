package main

import (
	"fmt"
)

type sumaElements struct {
	a int
	b int
}

func sumaProceso(sumasElementos chan sumaElements, resultados chan int, quit chan int) {
	for {
		select {
		case sumaARealizar := <-sumasElementos:
			resultados <- sumaARealizar.a + sumaARealizar.b
		case <-quit:
			fmt.Println("Go rutina parada")
			return
		}
	}
}

func main() {
	/*
		|
		|   \
		|-> | // Enviamos un sumaElementos
		|   |
		|<- | // Nos retorna el valor
		|   | sumaProceso
		|
		Main
	*/

	canalElementos := make(chan sumaElements)
	canalResultados := make(chan int, 50)
	canalQuit := make(chan int)
	go sumaProceso(canalElementos, canalResultados, canalQuit)
	canalElementos <- sumaElements{1, 2}
	i := 0
FR:
	for {
		select {
		case resultado := <-canalResultados:
			i++
			fmt.Println(resultado)
			if i == 10000000 {
				canalQuit <- 50
				break FR
			} else {
				canalElementos <- sumaElements{i, 2}
			}
		}
	}

}
