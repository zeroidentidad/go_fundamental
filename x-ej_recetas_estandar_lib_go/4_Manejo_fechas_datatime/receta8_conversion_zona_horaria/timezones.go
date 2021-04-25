package main

import (
	"fmt"
	"time"
)

func main() {

	mx, err := time.LoadLocation("America/Mexico_City")
	if err != nil {
		panic(err)
	}

	t := time.Date(2020, 1, 1, 0, 0, 0, 0, mx)
	fmt.Printf("Tiempo original (America/Mexico_City): %v\n", t)

	eur, err := time.LoadLocation("Europe/Vienna")
	if err != nil {
		panic(err)
	}

	phx, err := time.LoadLocation("America/Phoenix")
	if err != nil {
		panic(err)
	}

	t2 := t.In(eur)
	fmt.Printf("Tiempo convertido (Europe/Vienna): %v\n", t2)

	t3 := t.In(phx)
	fmt.Printf("Tiempo convertido (America/Phoenix): %v\n", t3)

}
