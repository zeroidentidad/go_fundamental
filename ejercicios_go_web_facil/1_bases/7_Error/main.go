package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/home", http.StatusMovedPermanently)
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Msge de error", http.StatusBadRequest)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// visualizar desde ternimal con: curl -i http://localhost:8080/home
