package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	serveWeb()
}

func serveWeb() {
	r := mux.NewRouter()

	r.HandleFunc("/", serveContent)
	r.HandleFunc("/{mi_parametro}", serveContent)

	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}

func serveContent(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	myParameter := urlParams["mi_parametro"]

	if myParameter == "" {
		myParameter = "Inicio"
	}

	w.Write([]byte("Hola " + myParameter))
}
