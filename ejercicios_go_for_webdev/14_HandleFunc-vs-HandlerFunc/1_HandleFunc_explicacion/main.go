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

	http.HandleFunc("/perro", d)
	http.HandleFunc("/gato", c)

	http.ListenAndServe(":8080", nil)
}
