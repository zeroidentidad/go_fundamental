package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	// short declaration
	days := 7
	fmt.Printf("%T %[1]v \n", days)

	var pi = math.Pi
	fmt.Printf("%T %[1]v \n", pi)

	var pi32 float32 = math.Pi
	fmt.Printf("%T %[1]v \n", pi32)

	tercio := 1.0 / 3
	fmt.Printf("%T %[1]v \n", tercio)
	fmt.Printf("%.2f \n", tercio)
	fmt.Printf("%f \n", tercio)

	distancia := new(big.Int)
	distancia.SetString("2400000000000000000000000000000000000000000", 10)
	fmt.Println(distancia)
}
