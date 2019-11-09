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

	mux := http.NewServeMux()
	mux.Handle("/inicio", jesus)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}

/*
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
*/
