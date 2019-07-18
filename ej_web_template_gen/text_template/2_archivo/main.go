package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("template.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("error creando archivo", err)
	}

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
