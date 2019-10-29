package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Holas ...")
		http.Redirect(w, r, "/home", http.StatusMovedPermanently)
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Holas kepedos")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// visualizar desde ternimal con: curl -i http://localhost:8080
// referencia: https://golang.org/src/net/http/status.go
