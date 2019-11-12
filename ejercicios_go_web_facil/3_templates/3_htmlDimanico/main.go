package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Usuario struct {
	UserName string
	Edad     int
}

func main() {
	port := "8080"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		template, err := template.ParseFiles("html/index.html")

		if err != nil {
			panic(err)
		}

		usuario := Usuario{UserName: "Zero", Edad: 100}
		template.Execute(w, usuario)
	})

	fmt.Println("Listen puerto :", port)
	http.ListenAndServe(":"+port, nil)
}
