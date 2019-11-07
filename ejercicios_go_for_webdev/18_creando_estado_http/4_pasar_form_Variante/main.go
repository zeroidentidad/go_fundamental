package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type persona struct {
	Nombre    string
	Apellido  string
	Subscrito bool
}

func main() {
	http.HandleFunc("/", watahea)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func watahea(w http.ResponseWriter, req *http.Request) {

	n := req.FormValue("nombre")
	a := req.FormValue("apellido")
	s := req.FormValue("subscrito") == "on"

	err := tpl.ExecuteTemplate(w, "index.gohtml", persona{n, a, s})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
