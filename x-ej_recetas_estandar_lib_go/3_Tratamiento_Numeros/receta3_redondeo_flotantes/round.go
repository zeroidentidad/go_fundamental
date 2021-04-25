package main

import (
	"fmt"
	"math"
)

var valA float64 = 3.55554444

func main() {

	// Mala suposición en el redondeo
	// el número enviándolo a
	// entero.
	intVal := int(valA)
	fmt.Printf("Mal redondeo al castear a int: %v\n", intVal)

	fRound := Roundf(valA)
	fmt.Printf("Redondeo por función personalizada: %v\n", fRound)

	mRound := math.Round(valA)
	fmt.Printf("Redondeo por paquete libreria Go: %v\n", mRound)

	mtRound := math.RoundToEven(valA)
	fmt.Printf("Redondeo por paquete libreria Go a par cercano: %v\n", mtRound)

}

// Roundf devuelve el entero más cercano.
func Roundf(x float64) float64 {
	t := math.Trunc(x)
	if math.Abs(x-t) >= 0.5 {
		return t + math.Copysign(1, x)
	}
	return t
}
