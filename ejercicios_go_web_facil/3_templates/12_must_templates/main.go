package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	port := "8080"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		template := template.Must(template.New("index.html").ParseFiles("html/index.html", "html/footer.html", "html/header.html"))
		// Como metodo

		template.Execute(w, nil)
	})

	fmt.Println("Listen puerto :", port)
	http.ListenAndServe(":"+port, nil)
}
