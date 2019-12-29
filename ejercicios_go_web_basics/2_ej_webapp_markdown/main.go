package main

import (
	"log"
	"net/http"
	"os"

	"github.com/russross/blackfriday"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))
	log.Println("Server iniciado")
	http.ListenAndServe(":"+port, nil)

}

// GenerateMarkdown con blackfriday lib
func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}

/*
V1:
Para uso básico, ingresar entrada en un segmento de bytes y llamar:

salida: = blackfriday.MarkdownBasic (entrada)

Esto lo hace sin extensiones habilitadas.

Para obtener un conjunto de funciones más útil, usar en su lugar:

salida: = blackfriday.MarkdownCommon (input)
*/
