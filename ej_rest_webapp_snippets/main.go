package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	const port = "4000"
	log.Println("Server start on port:", port)
	err := http.ListenAndServe(":"+port, mux)
	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home Snippets"))
}
