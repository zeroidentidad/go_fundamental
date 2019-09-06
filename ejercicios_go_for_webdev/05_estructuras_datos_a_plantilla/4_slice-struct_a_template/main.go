package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sabio struct {
	Nombre string
	Lema   string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	buddha := sabio{
		Nombre: "Buddha",
		Lema:   "La creencia de no creer",
	}

	gandhi := sabio{
		Nombre: "Gandhi",
		Lema:   "Sé el cambio",
	}

	mlk := sabio{
		Nombre: "Martin Luther King",
		Lema:   "El odio nunca cesa con el odio, pero solo con el amor se cura.",
	}

	jesus := sabio{
		Nombre: "Jesus",
		Lema:   "Amar todo",
	}

	muhammad := sabio{
		Nombre: "Muhammad",
		Lema:   "Vencer el mal con el bien es bueno, resistir el mal con el mal es malo.",
	}

	sabios := []sabio{buddha, gandhi, mlk, jesus, muhammad}

	err := tpl.Execute(os.Stdout, sabios)
	if err != nil {
		log.Fatalln(err)
	}
}
