package main

import "fmt"

func envoltura() func() int { // retornar una funcion con dato de tipo entero
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	incremento := envoltura()
	fmt.Println(incremento())
	fmt.Println(incremento())
}

/*
El cierre nos ayuda a limitar el alcance de las variables utilizadas por múltiples funciones
sin cierre, para que dos o más funciones tengan acceso a la misma variable,
esa variable tendría que estar en alcance del paquete
*/
