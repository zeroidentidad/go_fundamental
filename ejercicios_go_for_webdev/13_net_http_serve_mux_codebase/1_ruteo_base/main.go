package main

import (
	"io"
	"net/http"
)

type numx int

func (m numx) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/perro":
		io.WriteString(w, "wuau wuau... dijo el perro")
	case "/gato":
		io.WriteString(w, "miau miau... digo el gato")
	}
}

func main() {
	var d numx
	http.ListenAndServe(":8080", d)
}
