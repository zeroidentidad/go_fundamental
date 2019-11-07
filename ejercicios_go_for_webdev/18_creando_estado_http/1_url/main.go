package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", watahea)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func watahea(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	fmt.Fprintln(w, "Hacer mi busqueda:", v)
}

// visitar:
// http://localhost:8080/?q=perrosdelmal
