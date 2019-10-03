package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "Te escuche decir: %s\n", ln)
	}
	defer conn.Close()

	// nunca llegamos aquí
	// tenemos un stream de conexión abierta
	// ¿cómo sabe el reader anterior cuándo debe terminar?
	fmt.Println("Codigo extra aqui.")
}
