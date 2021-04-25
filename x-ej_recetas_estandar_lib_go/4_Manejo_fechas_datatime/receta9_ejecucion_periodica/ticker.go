package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c)

	ticker := time.NewTicker(time.Second)
	stop := make(chan bool)

	go func() {
		defer func() { stop <- true }()
		count := 0
		for {
			select {
			case <-ticker.C:
				count++
				fmt.Println("Tick ", count, time.Now())
			case <-stop:
				fmt.Println("\nGoroutina cerrandose")
				return
			}
		}
	}()

	// Bloquear hasta que se reciba la señal
	<-c
	ticker.Stop()

	// detener la goroutine
	stop <- true
	// espera hasta que
	<-stop
	fmt.Println("aplicacion detenida")
}
