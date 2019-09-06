package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	//sabios := map[string]string{"India":"Gandhi", "America":"MLK", "Meditar":"Buddha", "Amor":"Jesus", "Profeta":"Muhammad"}

	sabios := map[string]string{
		"India":   "Gandhi",
		"America": "MLK",
		"Meditar": "Buddha",
		"Amor":    "Jesus",
		"Profeta": "Muhammad"}

	err := tpl.Execute(os.Stdout, sabios)
	if err != nil {
		log.Fatalln(err)
	}
}
