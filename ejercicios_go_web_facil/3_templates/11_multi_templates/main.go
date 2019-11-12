package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type Curso struct {
	Nombre   string
	Duracion int
}

type Usuario struct {
	UserName string
	Edad     int
	Activo   bool
	Admin    bool
	Tags     []string
	Cursos   []Curso
}

func (este Usuario) TienePermisosAdmin(llave string) bool {
	return este.Activo && este.Admin && strings.ToLower(llave) == "si"
}

func hola() string {
	return "Hola :/ meh func"
}

func suma(v1, v2 int) int {
	return v1 + v2
}

func main() {
	port := "8080"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		funcionesmap := template.FuncMap{
			"hola": hola,
			"suma": suma,
		}

		//template, err := template.ParseFiles("html/index.html") // Como funcion

		template, err := template.New("index.html").Funcs(funcionesmap).ParseFiles("html/index.html", "html/footer.html", "html/header.html") // Como metodo

		if err != nil {
			panic(err)
		}

		tags := []string{"Go", "Python", "JavaScript", "C", "PHP"}

		cursos := []Curso{Curso{"Go", 24}, Curso{"JavaScript", 18}, Curso{"PHP", 12}}

		usuario := Usuario{
			UserName: "Zero",
			Edad:     100,
			Activo:   true,
			Admin:    true,
			Tags:     tags,
			Cursos:   cursos,
		}

		template.Execute(w, usuario)
	})

	fmt.Println("Listen puerto :", port)
	http.ListenAndServe(":"+port, nil)
}
