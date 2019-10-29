package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Nombre", "Valor ZeroI")
		w.Header().Add("Nombre2", "Valor ZeroI2")
		fmt.Fprintf(w, "Holas ...")
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Holas kepedos")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// visualizar desde ternimal con: curl -i http://localhost:8080
