package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Se crea el canal donde se recibirá,
	// se enviaría la señal. La Notify (func estandar lib)
	// no se bloqueará cuando la señal se envía
	// y el canal aun no este listo.
	// Entonces mejor se crea un canal almacenado en búfer.
	sChan := make(chan os.Signal, 1)

	// Notify atrapará las
	// señales dadas y enviará
	// el valor de os.Signal
	// a través de sChan
	signal.Notify(sChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGKILL)

	//Se crea un canal para esperar hasta que se maneje la señal.
	exitChan := make(chan int)
	go func() { // go rutina de cambio de señales
		signal := <-sChan
		switch signal {
		case syscall.SIGHUP:
			fmt.Println("La terminal de llamada ha sido cerrada.")
			exitChan <- 0

		case syscall.SIGINT:
			fmt.Println("El proceso ha sido interrumpido por CTRL+C")
			exitChan <- 1

		case syscall.SIGTERM:
			fmt.Println("kill SIGTERM se ejecutó para el proceso")
			exitChan <- 1

		case syscall.SIGKILL:
			fmt.Println("SIGKILL handler")
			exitChan <- 1

		case syscall.SIGQUIT:
			fmt.Println("kill SIGQUIT se ejecutó para el proceso")
			exitChan <- 1
		}
	}()

	code := <-exitChan
	os.Exit(code)
}
