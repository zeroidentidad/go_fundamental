package main

import (
	"log"
	"os"
	"text/template"
)

type curso struct {
	Numero, Nombre, Unidades string
}

type semestre struct {
	Temporada string
	Cursos    []curso
}

type anio struct {
	Otonio, Primavera, Verano semestre
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	y := anio{
		Otonio: semestre{
			Temporada: "Otonio",
			Cursos: []curso{
				curso{"CSCI-40", "Intro to Programming in Go", "4"},
				curso{"CSCI-130", "Intro to Web Programming with Go", "4"},
				curso{"CSCI-140", "Mobile Apps Using Go", "4"},
			},
		},
		Primavera: semestre{
			Temporada: "Primavera",
			Cursos: []curso{
				curso{"CSCI-50", "Advanced Go", "5"},
				curso{"CSCI-190", "Advanced Web Programming with Go", "5"},
				curso{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatalln(err)
	}
}
