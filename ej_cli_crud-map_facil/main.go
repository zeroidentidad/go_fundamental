package main

import (
	"fmt"
	"log"
)

func main() {
	var opcion string
	fmt.Println("A) Crear")
	fmt.Println("B) Listar")
	fmt.Println("C) Actualizar")
	fmt.Println("D) Eliminar")
	fmt.Scanln(&opcion)

	switch opcion {
	case "a", "crear":
		crearUsuario()

	case "b", "listar":
		listarUsuarios()

	case "c", "actualizar":
		actualizarUsuario()

	case "d", "eliminar":
		eliminarUsuarios()

	default:
		log.Println("Opción no valida")
	}

}

func crearUsuario() {
	log.Println("crearUsuario")
}

func listarUsuarios() {
	log.Println("listarUsuarios")
}

func actualizarUsuario() {
	log.Println("actualizarUsuario")
}

func eliminarUsuarios() {
	log.Println("eliminarUsuarios")
}
