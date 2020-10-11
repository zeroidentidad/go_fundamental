package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	const port = "4000"
	log.Println("Server start on port:", port)
	err := http.ListenAndServe(":"+port, mux)
	log.Fatal(err)
}
