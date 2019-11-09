package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	nombre := vars["nombre"]
	id := vars["id"]

	fmt.Fprintf(w, "Parametros: "+nombre+" "+id)

	w.Write([]byte("\nGorilla!"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/usuarios/{nombre}/{id:[0-9]+}", YourHandler)

	log.Fatal(http.ListenAndServe(":8000", r))
}
