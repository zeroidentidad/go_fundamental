package main

import (
	"fmt"
	"time"
)

func main() {

	loc, err := time.LoadLocation("America/Mexico_City")
	if err != nil {
		panic(err)
	}

	t := time.Date(2020, 11, 29, 21, 0, 0, 0, loc) //time.Local
	fmt.Printf("Extracción unidades de: %v\n", t)

	dOfMonth := t.Day()
	weekDay := t.Weekday()
	month := t.Month()

	fmt.Printf("El dia %d de %v es %v\n", dOfMonth, month, weekDay)

}
