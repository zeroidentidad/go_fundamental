package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hola)

	server := &http.Server{
		Addr:    ":8080",
		Handler: nil, // Se utiliza DefaultServeMux
	}

	log.Fatal(server.ListenAndServe())
}

func hola(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Holas ...")
}
