package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(aburriendo("Joe"), aburriendo("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Ambos son aburridos Me voy.")
}

func aburriendo(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

// FAN IN
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

/*
code source:
Rob Pike
https://talks.golang.org/2012/concurrency.slide#25
*/

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
