package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates = template.Must(template.New("W").ParseFiles("html/index.html", "html/footer.html", "html/header.html"))

func main() {
	port := "8080"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		err := templates.ExecuteTemplate(w, "index", nil)

		if err != nil {
			panic(err)
		}
	})

	fmt.Println("Listen puerto :", port)
	http.ListenAndServe(":"+port, nil)
}
