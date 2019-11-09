package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", hola2)
	mux.HandleFunc("/dos", hola3)

	http.HandleFunc("/", hola)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux, // Se utiliza DefaultServeMux
	}

	log.Fatal(server.ListenAndServe())
}

func hola(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Holas ...")
}

func hola2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "watahea ...")
}

func hola3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "watahea !!! ...")
}
