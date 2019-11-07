package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", watahea)
	http.HandleFunc("/redireccion", redireccion)
	http.HandleFunc("/redirijido", redirijido)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func watahea(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Método de solicitud en watahea: ", req.Method, "\n\n")
}

func redireccion(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Método de solicitud en redireccion:", req.Method)
	// procesar el envío del formulario aquí
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func redirijido(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Método de solicitud en redirijido:", req.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
