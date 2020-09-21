package main

import (
	"fmt"
	"math"
)

func main() {
	// formula: g(n,p)=q^(n-1)*p

	var n int
	var p float64
	//coeficiente de distribucion
	q := 0.3

	fmt.Println("Numero experimento:")
	fmt.Scanln(&n)
	fmt.Println("Probabilidad conocida:")
	fmt.Scanln(&p)

	result := math.Pow(q, float64(n-1)) * p

	fmt.Printf("%.5f", result)
}
