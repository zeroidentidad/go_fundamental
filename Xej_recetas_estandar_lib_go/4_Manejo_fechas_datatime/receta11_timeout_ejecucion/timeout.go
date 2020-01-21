package main

import (
	"fmt"
	"time"
)

func main() {

	to := time.After(4 * time.Second)
	list := make([]string, 0)
	done := make(chan bool, 1)

	fmt.Println("Iniciando insercion de elementos")
	go func() {
		defer fmt.Println("Saliendo de la gorutina")
		for {
			select {
			case <-to:
				fmt.Println("Se acabó el tiempo")
				done <- true
				return
			default:
				list = append(list, time.Now().String())
			}
		}
	}()

	<-done
	fmt.Printf("Gestionado para insertar %d elementos\n", len(list))

}
