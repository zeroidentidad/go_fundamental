// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 254.
//!+

// Chat is a server that lets clients chat with each other.

// Hay cuatro tipos de gorutina en este programa. Hay una instancia de cada una de las rutinas mainy broadcaster, y para cada conexión de cliente hay una handleConny una clientWritergorutina. La emisora ​​es un buen ejemplo de cómo selectse utiliza, ya que tiene que responder a tres tipos diferentes de mensajes.
//este programa es el servidor al que se conectan los clientes. cuando se conecta un cliente se escribe en el su direccion ip. si escribe un mensaje aparece tambien en el. Si conectamos otro nuevo cliente aparecera su direccion ip en los clientes ya conectados y en el mismo. Si con varios clientes escribimos un mensaje este mensaje aparecera en ambos.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

//!+broadcaster
type client struct { //creamos un tipo struct con el nombre del cliente y el canal
	exitmsg    chan<- string // an outgoing message channel
	nameclient string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() { //la funcion broadcaster escucha en el canal entering los clientes que entran y leaving borra los clientes que salen. Si entra un cliente actualiza clients añadiendo el nuevo cliente y si sale un cliente actualiza clients eliminando este cliente. Si sale un cliente cierra su canal de mensajes ya que no va a mandar mas.
	clients := make(map[client]bool) // all connected clients //clients registra el conjunto actual de clientes conectados
	for {
		select {
		case msg := <-messages: //tambien escucha eventos aqui. Recibe los mensajes en el messages.
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients { //los mensajes que recibe los manda a todos los clientes
				cli.exitmsg <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			cli.exitmsg <- "listclient:"
			for i := range clients { //bucle para escribir la lista de clientes conectados
				cli.exitmsg <- i.nameclient
			}

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.exitmsg)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn, nombre string) { //abre un canal de mensajes para el nuevo cliente conectado y avisa a los demas clientes de que este cliente se acaba de conectar a traves de entering
	ch := make(chan string /*, 1*/) // outgoing client messages
	go clientWriter(conn, ch)       //handleconn crea la gorutina clientwritter para cda cliente
	//in := make(chan string,1)
	//go clientReader(conn,in)

	/*ch <- "Enter your name:"
	sc := bufio.NewScanner(os.Stdin)
	_ = sc.Scan()

	who := sc.Text()*/
	//who := conn.RemoteAddr().String() //devuelve mi direccion ip
	cli := client{ch, nombre}
	ch <- "You are " + cli.nameclient
	messages <- cli.nameclient + " has arrived" //este mensaje se manda a los demas clientes ya conectados a traves de entering
	entering <- cli                             //cli

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		messages <- cli.nameclient + ": " + input.Text()
	}

	leaving <- cli //si un cliente se sale manda un mensaje a los demas clientes de que ha abandonado el chat y lo manda por leaving
	messages <- cli.nameclient + " has left"
	conn.Close() //cierra la conexion del cliente que se sale
}

func clientWriter(conn net.Conn, ch <-chan string) { //es para escribir el mensaje que escribe un cliente a el mismo y a todos los demas clientes
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		var nombre string
		fmt.Print("Nombre: ")
		fmt.Scanln(&nombre)
		go handleConn(conn, nombre)
	}
}

//!-main
