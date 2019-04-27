//Crear programa que declare 2 funciones anónimas.
// La primera va a contar del 100 al 0 .
// Mostrar con un identificador único para cada goroutine 
// Crear goruoutines de estas funciones y no permitir que main() regrese hasta que todas las goroutines se completen.

// Correr en paralelo

package main

import (
	"fmt"
	"runtime"
	"sync")

	// init se ejecuta antes de main()
	func init() {
		// Alojando/reservando (allocate) 1 procesador lógico para que lo use el scheduler.
		runtime.GOMAXPROCS(1)
	}

	// main es la entrada a nuestros programas

	func main() {

		// Declarar un wait group (grupo de espera) para iniciar el conteo en las 2 goroutines.

		var grupowait sync.WaitGroup
		grupowait.Add(2)

		fmt.Println("Iniciar Goroutines")

		// Declarar una función anónima y crear una goroutine

		go func() {
			//Cuenta regresiva del 100 al 0
			for contador := 100; contador >= 0; contador-- {
				fmt.Printf("[A:%d]\n", contador)
			}

			//Avisarle a main que ya terminamos 
			grupowait.Done()
		}()

		// Declarar una función anónima y crear una goroutine

		go func(){
			//Contar del 0 al 100
			for contador := 0; contador<=100; contador++{
				fmt.Printf("[B:%d]\n", contador)
			}

			//Decirle a main que ya terminamos
			grupowait.Done()
		}()

		// Esperar a que terminen las goroutines
		fmt.Println("Esperando a que terminen")
		grupowait.Wait()

		fmt.Println("\nFinalizar el programa")

	}
