package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "perro perro perro")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "gato gato gato")
}

func main() {

	http.Handle("/perro", http.HandlerFunc(d))
	http.Handle("/gato", http.HandlerFunc(c))

	http.ListenAndServe(":8080", nil)
}
