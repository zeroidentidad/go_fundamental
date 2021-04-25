package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	t := time.NewTimer(15 * time.Second)
	fmt.Printf("Comienza espera: %v\n", time.Now().Format(time.UnixDate))
	<-t.C
	fmt.Printf("Código ejecutado: %v\n", time.Now().Format(time.UnixDate))

	wg := &sync.WaitGroup{}
	wg.Add(1)

	fmt.Printf("Comienza espera para AfterFunc: %v\n", time.Now().Format(time.UnixDate))

	time.AfterFunc(15*time.Second, func() {
		fmt.Printf("Código ejecutado para AfterFunc: %v\n", time.Now().Format(time.UnixDate))
		wg.Done()
	})

	wg.Wait()

	fmt.Printf("Esperando en time.After: %v\n", time.Now().Format(time.UnixDate))
	<-time.After(15 * time.Second)
	fmt.Printf("Código reanudado: %v\n", time.Now().Format(time.UnixDate))

}
