package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Comment struct {
	Title       string
	Content     string
	ContentHTML template.HTML
}

var templates = template.Must(template.New("W").ParseGlob("html/**/*.html"))

var errorx = template.Must(template.ParseFiles("html/error.html"))

func handleError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	errorx.Execute(w, nil)
}

func renderTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	w.Header().Set("Content-Type", "text/html")

	err := templates.ExecuteTemplate(w, name, data)

	if err != nil {
		handleError(w, http.StatusInternalServerError)
	}
}

func comment(w http.ResponseWriter, r *http.Request) {
	title := "Ejemplo."
	content := "contenido del comentario <b>test safe</b>"
	contenthtml := template.HTML("contenido del comentario <b>test safe 2</b>")

	comment := Comment{
		Title:       title,
		Content:     content,
		ContentHTML: contenthtml,
	}

	renderTemplate(w, r, "comments/comment", comment)
}

func main() {
	port := "8080"

	mux := http.NewServeMux()

	mux.HandleFunc("/comment/", comment)

	fmt.Println("Listen puerto :", port)
	http.ListenAndServe(":"+port, mux)
}
