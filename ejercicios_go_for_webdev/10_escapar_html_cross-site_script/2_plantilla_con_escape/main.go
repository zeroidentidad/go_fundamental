package main

import (
	"html/template"
	"log"
	"os"
)

type Page struct {
	Title   string
	Heading string
	Input   string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	home := Page{
		Title:   "Escaped",
		Heading: "Lo peligroso se escapa con html/template",
		Input:   `<script>alert("js desde go, me muestro como texto html");</script>`,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", home)
	if err != nil {
		log.Fatalln(err)
	}

	/* go run main.go > index.html */
}
