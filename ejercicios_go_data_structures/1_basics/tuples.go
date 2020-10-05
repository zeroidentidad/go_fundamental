package main

import (
	"fmt"
)

// obtener la serie de potencias del entero "a" y devolver la tupla
// del cuadrado de "a" y el cubo de "a"

func powerSeries(a int) (int, int) { // forma 1

	return a * a, a * a * a

}

func powerSeriesN(a int) (square int, cube int) { // forma 2

	square = a * a

	cube = square * a

	return

}

func powerSeriesE(a int) (int, int, error) { // forma 3

	var square int = a * a

	var cube int = square * a

	return square, cube, nil

}

func main() {

	var square int
	var cube int
	square, cube = powerSeries(3)

	fmt.Println("Cuadrado ", square, "Cubo ", cube)

	fmt.Println(powerSeriesN(4))

	fmt.Println(powerSeriesE(5))
}
