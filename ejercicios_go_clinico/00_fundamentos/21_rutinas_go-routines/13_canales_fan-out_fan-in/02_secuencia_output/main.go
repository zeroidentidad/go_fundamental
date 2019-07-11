package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3)

	// FAN salida
	// Distribuye la secuencia de trababajo entre dos gorsalidaines que ambos leen desde adentro.
	c1 := sq(in)
	c2 := sq(in)

	// FAN IN
	// Consume la salida combinada de múltiples canales.
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 entonces 9, o 9 entonces 4
	}
}

func gen(nums ...int) chan int {
	fmt.Printf("TIPO DE NUMS %T\n", nums) // solo para tu información

	salida := make(chan int)
	go func() {
		for _, n := range nums {
			salida <- n
		}
		close(salida)
	}()
	return salida
}

func sq(in chan int) chan int {
	salida := make(chan int)
	go func() {
		for n := range in {
			salida <- n * n
		}
		close(salida)
	}()
	return salida
}

func merge(cs ...chan int) chan int {
	fmt.Printf("TYPE OF CS: %T\n", cs) // just FYI

	salida := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch chan int) {
			for n := range ch {
				salida <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(salida)
	}()

	return salida
}

/*
FAN OUT
Múltiples funciones de lectura desde el mismo canal hasta que ese canal se cierre.

FAN IN
Una función puede leer desde varias entradas y continuar hasta que todas estén cerradas por
multiplexar los canales de entrada en un solo canal que se cierra cuando todas las entradas están cerradas.

PATRON
Hay un patrón en nuestras funciones de pipeline:
- las etapas cierran sus canales de salida cuando se realizan todas las operaciones de envío.
- las etapas siguen recibiendo valores de los canales entrantes hasta que esos canales se cierran.

source:
https://blog.golang.org/pipelines
*/

/*
DESAFIO # 1
Cuando se sabe cómo hacer fan out / fan in, pero ¿sabemos POR QUÉ?
¿Por qué querríamos hacer fan out / fan in?
*/
