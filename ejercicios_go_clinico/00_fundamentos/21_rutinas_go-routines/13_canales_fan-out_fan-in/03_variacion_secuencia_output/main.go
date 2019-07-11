package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3)

	// Distribuye la secuencua de trabajo entre dos goroutines que ambos leen desde adentro.
	c1 := sq(in)
	c2 := sq(in)

	// Consume la salida combinada de c1 y c2.
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 entonces 9, o 9 entonces 4
	}
}

func gen(nums ...int) <-chan int {
	salida := make(chan int)
	go func() {
		for _, n := range nums {
			salida <- n
		}
		close(salida)
	}()
	return salida
}

func sq(in <-chan int) <-chan int {
	salida := make(chan int)
	go func() {
		for n := range in {
			salida <- n * n
		}
		close(salida)
	}()
	return salida
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	salida := make(chan int)

	// Comienza una salida de goroutine para cada canal de entrada en cs.
	// la salida copia los valores de c a salida hasta que c se cierra, luego llama a wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			salida <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Inicia una goroutine para cerrar la salida una vez que todos los goroutines de output son
	// hechos. Esto debe comenzar después de la llamada wg.Add
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

MODELO
Hay un patrón en nuestras funciones de pipeline:
- las etapas cierran sus canales de salida cuando se realizan todas las operaciones de envío.
- las etapas siguen recibiendo valores de los canales entrantes hasta que esos canales se cierran.

source:
https://blog.golang.org/pipelines
*/
