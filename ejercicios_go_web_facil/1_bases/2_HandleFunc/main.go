package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//w.Write([]byte{1})
		fmt.Fprintf(w, "Holas ...")
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Holas kepedos")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
