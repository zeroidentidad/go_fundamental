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

	t := time.Date(2020, 1, 1, 0, 0, 0, 0, loc) //time.Local
	t2 := time.Date(2010, 1, 3, 0, 0, 0, 0, loc)
	fmt.Printf("Primera fecha predeterminada es %v\n", t)
	fmt.Printf("Segunda fecha predeterminada es %v\n", t2)

	dur := t2.Sub(t)
	fmt.Printf("Duración entre var t y t2 es %v\n", dur)

	dur = time.Since(t)
	fmt.Printf("Duración entre tiempo actual y var t es %v\n", dur)

	dur = time.Until(t)
	fmt.Printf("Duración entre var t y tiempo actual es %v\n", dur)

}
