package main

import (
    "fmt"
    "sync"
)

const (
    goroutines = 20
    capacidad = 4
)

// grupowait se usa para esperar que el programa termine
var grupowait sync.WaitGroup

// recursos es un buffered channel para manejar los strings
var recursos = make(chan string, capacidad)

// main entrada a programa
func main() {
    // agregar el número de goroutines al waitgroup
    grupowait.Add(goroutines)

    // lanzar las goroutines necesarias para manejar el trabajo.
    // Asegurar de poner la llamada para saber que las gorutines terminaron
    for rutina := 1; rutina <= goroutines; rutina++ {
        go func(rutina int) {
            worker(rutina)
            grupowait.Done()
        }(rutina)
    }
    // Añadimos los strings.
    for s := 'A'; s < 'A' + capacidad; s++ {
        recursos <- string(s)
    }
    grupowait.Wait()
}

// lanzamos worker como una goroutine que procesa el trabajo del buffer channel
func worker(worker int){
    // Recibir un string del channel
    valor := <- recursos
    //  Imprimir el valor
    fmt.Printf("worker: %d : %s\n ", worker, valor)
    // Poner el string de regreso
    recursos <- valor
}