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

type doubleZero struct {
	person
	LicenciaMatar bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	p1 := doubleZero{
		person{
			Nombre: "Jesus Ferrer",
			Edad:   100,
		},
		true,
	}

	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}

}
