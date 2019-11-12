package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Usuario struct {
	UserName string
	Edad     int
	Activo   bool
	Admin    bool
}

func main() {
	port := "8080"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		template, err := template.ParseFiles("html/index.html")

		if err != nil {
			panic(err)
		}

		usuario := Usuario{
			UserName: "Zero",
			Edad:     100,
			Activo:   false,
			Admin:    true,
		}

		template.Execute(w, usuario)
	})

	fmt.Println("Listen puerto :", port)
	http.ListenAndServe(":"+port, nil)
}
