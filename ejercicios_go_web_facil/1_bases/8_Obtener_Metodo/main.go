package main

import (
	"fmt"
	"log"
	"net/http"
)

//GET, POST, PUT, DELETE

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Metodo usado: ", r.Method)

		switch r.Method {
		case "GET":
			fmt.Fprintf(w, "Hola, usando metodo: GET")
		case "POST":
			fmt.Fprintf(w, "Hola, usando metodo: POST")
		case "PUT":
			fmt.Fprintf(w, "Hola, usando metodo: PUT")
		case "PATCH":
			fmt.Fprintf(w, "Hola, usando metodo: PATCH")
		case "DELETE":
			fmt.Fprintf(w, "Hola, usando metodo: DELETE")
		default:
			http.Error(w, "No valido", http.StatusBadRequest)
		}

	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// test metodos:
// curl http://localhost:8080 -X GET
// curl http://localhost:8080 -X POST -i
