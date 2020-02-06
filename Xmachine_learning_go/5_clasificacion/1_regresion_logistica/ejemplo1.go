package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(logistic(1.0))
}

// Logistic implementa la función logistic,
// que se utiliza en la regresión logística.
func logistic(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}
