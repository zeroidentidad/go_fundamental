package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var writer *os.File

func main() {

	// El archivo se abre como
	// un archivo de registro para escribir (Log).
	// De esta manera se representan los recursos asignados.
	var err error
	writer, err = os.OpenFile(fmt.Sprintf("prueba_%d.log", time.Now().Unix()), os.O_RDWR|os.O_CREATE, os.ModePerm) // %d format flag numerica
	if err != nil {
		panic(err)
	}

	// El código se queda ejecutando en gorutina.
	// Entonces, en caso de que el programa sea
	// terminado desde afuera, necesitamos
	// deja que la gorutina sepa a través de closeChan
	closeChan := make(chan bool)
	go func() {
		for {
			time.Sleep(time.Second)
			select {
			case <-closeChan:
				log.Println("Goroutina cerrandose")
				return
			default:
				log.Println("Escribiendo log")
				io.WriteString(writer, fmt.Sprintf("Acceso logging %s\n", time.Now().String()))
			}

		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGINT)

	// Esto bloquea la lectura de
	// sigChan donde la función Notify envía
	// la señal.
	<-sigChan

	// Después de recibir la señal
	// todo el código detrás de la lectura del canal podría ser
	// considerado como una limpieza de recursos
	close(closeChan)
	releaseAllResources()
	fmt.Println("La aplicación se cerró con sutilmente.")
}

func releaseAllResources() {
	io.WriteString(writer, "Aplicación liberando todos los recursos\n")
	writer.Close()
}
