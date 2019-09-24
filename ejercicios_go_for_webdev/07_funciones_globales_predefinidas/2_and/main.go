package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type user struct {
	Nombre string
	Lema   string
	Admin  bool
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	u1 := user{
		Nombre: "Buddha",
		Lema:   "La creencia de no creer",
		Admin:  false,
	}

	u2 := user{
		Nombre: "Gandhi",
		Lema:   "Sé el cambio",
		Admin:  true,
	}

	u3 := user{
		Nombre: "",
		Lema:   "Nadie",
		Admin:  true,
	}

	users := []user{u1, u2, u3}

	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}
}
