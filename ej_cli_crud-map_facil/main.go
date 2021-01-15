package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Println("A) Crear")
		fmt.Println("B) Listar")
		fmt.Println("C) Actualizar")
		fmt.Println("D) Eliminar")
		fmt.Println("Ingresar opcion:")
		opcion := readLine(reader)

		if opcion == "quit" || opcion == "q" {
			break
		}

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

	fmt.Println("Adios...")

}

func readLine(reader *bufio.Reader) string {
	if opcion, err := reader.ReadString('\n'); err != nil {
		panic("No se pudo obtener valor!")
	} else {
		return strings.TrimSuffix(opcion, "\n")
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
