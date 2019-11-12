package main

import (
	"fmt"
	"html/template"
	"net/http"
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

func main() {
	port := "8080"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		template, err := template.ParseFiles("html/index.html")

		if err != nil {
			panic(err)
		}

		tags := []string{"Go", "Python", "JavaScript", "C", "PHP"}

		cursos := []Curso{Curso{"Go", 24}, Curso{"JavaScript", 18}, Curso{"PHP", 12}}

		usuario := Usuario{
			UserName: "Zero",
			Edad:     100,
			Activo:   false,
			Admin:    true,
			Tags:     tags,
			Cursos:   cursos,
		}

		template.Execute(w, usuario)
	})

	fmt.Println("Listen puerto :", port)
	http.ListenAndServe(":"+port, nil)
}
