package main

import (
	"fmt"
	"log"
	"net/http"
)

type messageHandler struct {
	message string
}

func (m messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.message)
}

func main() {
	mux := http.NewServeMux()
	meg1 := &messageHandler{message: "<h1>Hola desde un handler</h1>"}
	meg2 := &messageHandler{message: "<h1>Desde otro</h1>"}
	meg3 := &messageHandler{message: "<h1>Alguna cosa aqui</h1>"}

	mux.Handle("/saludar", meg1)
	mux.Handle("/otro", meg2)
	mux.Handle("/cosa", meg3)

	log.Println("Server ejecutandose en :8080")
	log.Println(http.ListenAndServe(":8080", mux))
}
