package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Nombre string
	Edad   int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	p1 := person{
		Nombre: "Jesus Ferrer",
		Edad:   100,
	}

	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}

}
