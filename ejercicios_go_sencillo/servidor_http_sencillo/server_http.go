package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

func iniciarServidor(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Servidor creado con lenguaje Go funcionando!, Hola, 你好, 안녕하세요., привет ")

}

func main() {
	http.HandleFunc("/", iniciarServidor)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error al iniciar http server : ", err)
		return

	}

}
