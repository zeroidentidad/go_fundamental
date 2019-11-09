package main

import (
	"fmt"
	"log"
	"net/http"
)

type Usuario struct {
	nombre string
}

func (este *Usuario) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whats up "+este.nombre)
}

func main() {

	jesus := &Usuario{nombre: "Jesus :v"}

	server := &http.Server{
		Addr:    ":8080",
		Handler: jesus,
	}

	log.Fatal(server.ListenAndServe())
}

/*
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
*/
