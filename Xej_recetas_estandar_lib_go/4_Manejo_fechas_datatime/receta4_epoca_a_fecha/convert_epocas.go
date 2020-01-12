package main

import (
	"fmt"
	"time"
)

func main() {

	// Establecer la época de int64
	t := time.Unix(0, 0)
	fmt.Println(t)

	// Obtener la época de la instancia Time
	epoca := t.Unix()
	fmt.Println(epoca)

	// Tiempo actual
	epocaNow := time.Now().Unix()
	fmt.Printf("Época en segundos: %d\n", epocaNow)

	epocaNano := time.Now().UnixNano()
	fmt.Printf("Época en nano-segundos: %d\n", epocaNano)

}
