package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "mi-cookie",
		Value: "some value",
		Path:  "/",
	})
	fmt.Fprintln(w, "COOKIE ESCRITA - COMPROBAR EN NAVEGADOR")
	fmt.Fprintln(w, "en chrome ir a: dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {

	c1, err := req.Cookie("mi-cookie")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "- COOKIE #1:", c1)
	}

	c2, err := req.Cookie("general")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "- COOKIE #2:", c2)
	}

	c3, err := req.Cookie("especifica")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "- COOKIE #3:", c3)
	}
}

func abundance(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "algun valor sobre algo general",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "especifica",
		Value: "algun valor sobre algo especifico",
	})

	fmt.Fprintln(w, "COOKIES ESCRITAS - COMPROBAR EN NAVEGADOR")
	fmt.Fprintln(w, "en chrome ir a: dev tools / application / cookies")
}
