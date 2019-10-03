package main

import (
	"fmt"
	"io"
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

		io.WriteString(conn, "\nHola desde el servidor TCP\n")
		fmt.Fprintln(conn, "Que onda?")
		fmt.Fprintf(conn, "%v", "Pues, nada tranquilon...")

		conn.Close()
	}
}

// linux test: desde win7 -> telnet 192.168.0.103 8080
