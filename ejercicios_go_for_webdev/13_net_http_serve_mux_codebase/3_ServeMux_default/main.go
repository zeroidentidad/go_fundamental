package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "wuau wuau... dijo el perro")
}

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "miau miau... digo el gato")
}

func main() {
	var d hotdog
	var c hotcat

	http.Handle("/perro", d)
	http.Handle("/gato", c)

	http.ListenAndServe(":8080", nil)
}
