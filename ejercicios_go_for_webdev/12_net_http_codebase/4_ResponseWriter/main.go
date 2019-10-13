package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Zero-Key", "esto es de Jesus Ferrer")
	w.Header().Set("HandShake", "saludo prros :v")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Cualquier código que quieras en esta función</h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
