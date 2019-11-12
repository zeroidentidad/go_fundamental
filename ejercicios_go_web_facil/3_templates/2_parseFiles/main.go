package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	port := "8080"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		template, err := template.ParseFiles("html/index.html")

		if err != nil {
			panic(err)
		}

		template.Execute(w, nil)
	})

	fmt.Println("Listen puerto :", port)
	http.ListenAndServe(":"+port, nil)
}
