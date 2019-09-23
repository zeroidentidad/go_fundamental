package main

import (
	"fmt"
)

//  - Por defecto tenemos 3 tipos de funciones.

//   - Funciones sin parámetros de entrada ni de salida
func main() {
	suma(2, 2)
	sumaRecortada(2, 2)
	fmt.Println(sumaRetorno(2, 3))
	fmt.Println(sumaCombinada(2, 3, 4))
	fmt.Println(sumaCombinadaNombrada(2, 3, 4))
}

//   - Funciones con parámetros de entrada, pero sin parámetros de salida
func suma(x int, y int) {
	fmt.Println(x + y)
}

// Podemos emplear Sugar Syntax para recortar las funciones agrupando los tipos de variables
// siempre que sean iguales.
func sumaRecortada(x, y int) {
	fmt.Println(x + y)
}

//   - Funciones con parámetros de entrada y de salida
func sumaRetorno(x int, y int) int {
	return x + y
}

// Podemos devolver más de un parámetro de salida
func sumaCombinada(x int, y int, z int) (int, int) {
	return x + y, x + z
}

// Podemos nombrar los parámetros para trabajar con ellos dentro de una función
// como si de un parámetro más se tratase
func sumaCombinadaNombrada(x int, y int, z int) (l int, j int) {
	l = x + y
	j = x + z
	return
}
