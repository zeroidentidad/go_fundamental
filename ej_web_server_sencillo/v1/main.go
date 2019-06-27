package main

import (
	"net/http"
)

func main() {

	// rutas
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contacto", contactHandler)

	// start serv
	http.ListenAndServe(":80", nil)
	// para poder usar el port 80 se necesita ejecucion como admin, ej.: sudo go run main.go
}

// funciones de manejo de rutas en ambito del paquete actual
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola mundo papu"))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pagina de contacto"))
}
