package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", watahea)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func watahea(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}
