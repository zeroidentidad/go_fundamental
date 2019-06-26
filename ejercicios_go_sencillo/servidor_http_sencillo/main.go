package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	connHost = "localhost"
	connPort = "8080"
)

func iniciarServidor(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Servidor iniciado!, Hola mundo :v ")

}

func main() {
	http.HandleFunc("/", iniciarServidor)
	err := http.ListenAndServe(connHost+":"+connPort, nil)
	if err != nil {
		log.Fatal("error al iniciar http server: ", err)
		return

	}

}
