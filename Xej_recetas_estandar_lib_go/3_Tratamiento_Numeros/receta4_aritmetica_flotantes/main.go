package main

import (
	"fmt"
	"math/big"
)

const PI = `3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862803482534211706798214808651328230664709384460955058223172535940812848111745028410270193852110555964462294895493038196`
const diametro = 3.0
const precision = 400

func main() {

	pi, _ := new(big.Float).SetPrec(precision).SetString(PI)
	d := new(big.Float).SetPrec(precision).SetFloat64(diametro)

	circunferencia := new(big.Float).Mul(pi, d)

	pi64, _ := pi.Float64()
	fmt.Println("_____________")
	fmt.Printf("Circunferencia big.Float = %.100f\n", circunferencia)
	fmt.Println("_____________")
	fmt.Printf("Circunferencia float64   = %.100f\n", pi64*diametro)

	sum := new(big.Float).Add(pi, pi)
	fmt.Println("_____________")
	fmt.Printf("Suma = %.100f\n", sum)

	diff := new(big.Float).Sub(pi, pi)
	fmt.Println("_____________")
	fmt.Printf("Diferencia = %.100f\n", diff)

	quo := new(big.Float).Quo(pi, pi)
	fmt.Println("_____________")
	fmt.Printf("Cociente = %.100f\n", quo)

}
