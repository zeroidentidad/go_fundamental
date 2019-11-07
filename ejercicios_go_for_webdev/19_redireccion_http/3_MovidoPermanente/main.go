package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", watahea)
	http.HandleFunc("/redireccion", redireccion)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func watahea(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Método de solicitud en watahea: ", req.Method, "\n\n")
}

func redireccion(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Método de solicitud en redireccion:", req.Method)
	http.Redirect(w, req, "/", http.StatusMovedPermanently)
}
