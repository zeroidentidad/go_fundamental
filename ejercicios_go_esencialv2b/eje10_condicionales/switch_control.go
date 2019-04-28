package main

import (
	"fmt"
	"time"
)

func main() {
	tiempo := time.Now()

	switch diaHoy := tiempo.Weekday(); diaHoy {
	case 0:
		fmt.Println("Es dia de descanso Domingo")
	case 1:
		fmt.Println("Gran inicio es Lunes")
	case 2:
		fmt.Println("Ok, es martes")
	default:
		fmt.Println("Es cualquier otro dia")

	}

}
