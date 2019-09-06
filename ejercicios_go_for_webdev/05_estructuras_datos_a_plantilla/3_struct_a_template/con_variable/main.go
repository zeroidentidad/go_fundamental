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

	err := tpl.Execute(os.Stdout, buddha)
	if err != nil {
		log.Fatalln(err)
	}
}
