
// declarar mapa de valores enteros con un string como llave

// llenar el mapa con 5 valores e iterar sobre el map para mostrar los pares llave/valor

package main

import "fmt"

func main(){

	// declaracion
	departamentos := make(map[string]int)

	// array de las claves
	deps:= [5] string {"Devs", "Marketing", "Ejecutivos", "Ventas", "Manteminiento"}

	// asignar valores en el mapa

	departamentos[deps[0]] = 25
	departamentos[deps[1]] = 40
	departamentos[deps[2]] = 3
	departamentos[deps[3]] = 50
	departamentos[deps[4]] = 6

	// deplegar con iteracion los clave-valor
	for clave, valor := range departamentos{
		fmt.Printf("Depto: %s - Personas: %d\n", clave, valor) 
		// %s placeholder de string, %d placeholder de numeros, \n salto de linea
	}

}