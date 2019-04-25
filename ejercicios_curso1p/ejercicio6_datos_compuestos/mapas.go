
// declarar mapa de valores enteros con un string como llave

// llenar el mapa con 5 valores e iterar sobre el map para mostrar los pares llave/valor

package main

import "fmt"

func main(){

	// declaracion
	departamentos := make(map[string]int)

	// asignar valores en el mapa

	departamentos["Devs"] = 25
	departamentos["Marketing"] = 40
	departamentos["Ejecutivos"] = 3
	departamentos["Ventas"] = 50
	departamentos["Manteminiento"] = 6

	// deplegar con iteracion los clave-valor
	for clave, valor := range departamentos{
		fmt.Printf("Depto: %s - Personas: %d\n", clave, valor) 
		// %s placeholder de string, %d placeholder de numeros, \n salto de linea
	}

}