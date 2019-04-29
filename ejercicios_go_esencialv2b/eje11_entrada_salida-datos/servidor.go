package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", saludo)
	http.ListenAndServe(":8000", nil)
}

func saludo(respuesta http.ResponseWriter, peticion *http.Request) {
	io.WriteString(respuesta, "Hola mundo web :v")
}
