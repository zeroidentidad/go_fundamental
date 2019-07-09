package main

import "fmt"

var x int

func incremento() int {
	x++
	return x
}

func main() {
	fmt.Println(incremento())
	fmt.Println(incremento())
}

/*
El cierre nos ayuda a limitar el alcance de las variables utilizadas por múltiples funciones
sin cierre, para que dos o más funciones tengan acceso a la misma variable,
esa variable tendría que estar en el alcance del paquete
*/
