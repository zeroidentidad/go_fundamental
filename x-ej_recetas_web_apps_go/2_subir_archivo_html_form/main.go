package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", fileHandler)

	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error al iniciar el servidor http: ", err)
		return
	}
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("archivo")
	if err != nil {
		log.Printf("error al obtener un archivo para la key de formulario proporcionada: ", err)
		return
	}
	defer file.Close()

	out, pathError := os.Create("/tmp/uploadedFile")
	if pathError != nil {
		log.Printf("error al crear un archivo para escribir: ", pathError)
		return
	}
	defer out.Close()

	_, copyFileError := io.Copy(out, file)

	if copyFileError != nil {
		log.Printf("se produjo un error al copiar el archivo: ", copyFileError)
	}

	fmt.Fprintf(w, "Archivo cargado exitosamente: "+header.Filename)
}

func index(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("templates/upload-file.html")
	parsedTemplate.Execute(w, nil)
}
