
// declarar struct para mantener info de un usuario (nombre, direccion, edad)

// Crear un valor e inicializar con valores
// Mostrar cada campo
// Declarar e inicializar un struct anónimo con los mismos 3 campos
// Mostrar el valor.

package main

import "fmt"

// usuario representa un usuario en el sistema

type usuario struct {
	nombre string
	direccion string
	edad int
}

func main(){

	jesus := usuario{
		nombre: "Jesus",
		direccion: "Tabasco",
		edad: 26,
	}

	//mostrar valores
	fmt.Println("Nombre: ", jesus.nombre)
	fmt.Println("Dirección: ", jesus.direccion)
	fmt.Println("Edad: ", jesus.edad)

	//declarar struct anonimo con los mismos 3 campos
	otro := struct{
		nombre string
		direccion string
		edad int
	} {
		nombre: "Otro",
		direccion: "Lugar x",
		edad: 22,
	}

	//mostrar valores
	fmt.Println("Nombre: ", otro.nombre)
	fmt.Println("Dirección: ", otro.direccion)
	fmt.Println("Edad: ", otro.edad)


}