package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type User struct {
	id       int
	username string
	email    string
	age      int
}

type ReadStd struct {
	Reader *bufio.Reader
}

var id int
var users map[int]User
var r ReadStd

func main() {

	users = make(map[int]User)
	for {
		fmt.Println("A) Crear")
		fmt.Println("B) Listar")
		fmt.Println("C) Actualizar")
		fmt.Println("D) Eliminar")
		fmt.Println("Ingresar opcion:")
		opcion := r.readLine()

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

	log.Println("Adios...")
}

func (r ReadStd) readLine() string {
	r.Reader = bufio.NewReader(os.Stdin)
	if opcion, err := r.Reader.ReadString('\n'); err != nil {
		panic("No se pudo obtener valor!")
	} else {
		return strings.TrimSuffix(opcion, "\n")
	}
}

func crearUsuario() {
	clearConsole()
	fmt.Println("Ingresa valor para username:")
	username := r.readLine()

	fmt.Println("Ingresa valor para email:")
	email := r.readLine()

	fmt.Println("Ingresa valor para edad:")
	age, err := strconv.Atoi(r.readLine())
	if err != nil {
		panic("Valor no valido")
	}

	id++
	user := User{
		id,
		username,
		email,
		age,
	}

	users[id] = user
}

func listarUsuarios() {
	clearConsole()
	for id, user := range users {
		fmt.Println(id, "-", user)
	}
}

func actualizarUsuario() {
	clearConsole()

	fmt.Println("Ingresar id de usuario a actualizar:")
	id, err := strconv.Atoi(r.readLine())
	if err != nil {
		panic("Valor no valido")
	}

	fmt.Println("Ingresa valor para username:")
	username := r.readLine()

	fmt.Println("Ingresa valor para email:")
	email := r.readLine()

	fmt.Println("Ingresa valor para edad:")
	age, err := strconv.Atoi(r.readLine())
	if err != nil {
		panic("Valor no valido")
	}

	/*user := User{
		id,
		username,
		email,
		age,
	}

	users[id] = user*/

	for _, user := range users {
		if user.id == id {
			if username != "" {
				user.username = username
			}
			if email != "" {
				user.email = email
			}
			if age != 0 {
				user.age = age
			}
			users[id] = user
		}
	}
}

func eliminarUsuarios() {
	clearConsole()
	fmt.Println("Ingresar id de usuario a eliminar:")
	id, err := strconv.Atoi(r.readLine())
	if err != nil {
		panic("Valor no valido")
	}

	if _, ok := users[id]; ok {
		delete(users, id)
	}

}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}
