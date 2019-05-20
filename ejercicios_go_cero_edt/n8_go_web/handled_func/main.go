package main

import (
	"fmt"
	"log"
	"net/http"
)

func messageHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola con HandlerFunc</h1>")
}

func messageHFCustom(message string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, message)
		},
	)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", messageHandlerFunc)

	mux.Handle("/saludar", messageHFCustom("<h1>Hola mundo</h1>"))
	mux.Handle("/despedirse", messageHFCustom("<h1>Adios mundo cruel</h1>"))

	log.Println("Server ejecutandose en :8080")
	log.Println(http.ListenAndServe(":8080", mux))
}
