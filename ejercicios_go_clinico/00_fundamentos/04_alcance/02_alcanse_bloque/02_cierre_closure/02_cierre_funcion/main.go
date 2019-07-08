package main

import "fmt"

var x = 0

func increment() int {
	x++
	return x
}

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
El cierre nos ayuda a limitar el alcance de las variables utilizadas por múltiples funciones
sin cierre, para que dos o más funciones tengan acceso a la misma variable,
esa variable tendría que ser en alcance del paquete
*/
