package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Usuario struct {
	UserName string
}

var templates = template.Must(template.New("W").ParseGlob("html/**/*.html"))

func renderTemplate(w http.ResponseWriter, name string, data interface{}) {
	w.Header().Set("Content-Type", "text/html")

	err := templates.ExecuteTemplate(w, name, data)

	if err != nil {
		http.Error(w, "No es posible retornar template", http.StatusInternalServerError)
	}
}

func main() {
	port := "8080"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index", nil)
	})

	http.HandleFunc("/usuario", func(w http.ResponseWriter, r *http.Request) {
		usuario := Usuario{"Jesus"}
		renderTemplate(w, "usuario", usuario)
	})

	fmt.Println("Listen puerto :", port)
	http.ListenAndServe(":"+port, nil)
}
