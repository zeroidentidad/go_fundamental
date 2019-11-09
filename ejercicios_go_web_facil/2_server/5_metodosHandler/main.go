package main

import (
	"log"
	"net/http"
)

func main() {

	redirect := http.RedirectHandler("https://google.com", http.StatusMovedPermanently)
	notfound := http.NotFoundHandler()

	mux := http.NewServeMux()

	mux.Handle("/redirect", redirect)
	mux.Handle("/notfound", notfound)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
