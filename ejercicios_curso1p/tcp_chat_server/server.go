package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// quien que emite los mensajes
type cliente chan<- string

// los mensajes
var (
	entrantes = make(chan cliente)
	salientes = make(chan cliente)
	mensajes  = make(chan string)
)

// funcion emisora/difusora
func broadcaster() {
	// todos los clientes conectados
	clientes := make(map[cliente]bool)

	for {
		select {
		case msg := <-mensajes:
			// Emitir mansaje entrante a todos los clientes conectados de channels
			for cli := range clientes {
				cli <- msg
			}
		case cli := <-entrantes:
			clientes[cli] = true
		case cli := <-salientes:
			delete(clientes, cli)
			close(cli)
		}
	}
}

func handlerConexion(conexion net.Conn) {
	canal := make(chan string)
	go clientWriter(conexion, canal)

	hostname := conexion.RemoteAddr().String()
	canal <- "Tú eres: " + hostname

	mensajes <- hostname + " se ha conectado"
	entrantes <- canal

	input := bufio.NewScanner(conexion)
	for input.Scan() {
		mensajes <- hostname + ": " + input.Text()
	}

	salientes <- canal
	mensajes <- hostname + " se ha ido"
	conexion.Close()
}

func clientWriter(conexion net.Conn, canal <-chan string) {
	for msg := range canal {
		fmt.Fprintln(conexion, msg)
	}
}

func main() {

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conexion, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handlerConexion(conexion)
	}

}
