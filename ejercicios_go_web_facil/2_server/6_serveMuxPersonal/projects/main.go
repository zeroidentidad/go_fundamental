package main

import (
	"fmt"
	"log"
	"net/http"

	"./mux"
)

type Usuario struct {
	nombre string
}

func (este *Usuario) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "whatahea "+este.nombre)
}

func main() {

	usuario := &Usuario{"Jesus"}

	mux := mux.CreateMux()
	mux.AddFunc("/holas", Hola)
	mux.AddHandle("/usuario", usuario)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func Hola(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "watahea from func anonima")
}
