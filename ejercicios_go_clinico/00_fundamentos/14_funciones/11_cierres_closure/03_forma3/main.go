package main

import "fmt"

func main() {
	x := 0
	incremento := func() int {
		x++
		return x
	}
	fmt.Println(incremento())
	fmt.Println(incremento())
}

/*
El cierre nos ayuda a limitar el alcance de las variables utilizadas por múltiples funciones
sin cierre, para que dos o más funciones tengan acceso a la misma variable,
esa variable tendría que estar en alcance del paquete

función anónima
una función sin nombre

expresion func
asignando una func a una variable
*/
