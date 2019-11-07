package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", watahea)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func watahea(res http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("mi-cookie")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "mi-cookie",
			Value: "0",
			Path:  "/",
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(res, cookie)

	io.WriteString(res, cookie.Value)

	fmt.Fprintln(res, ` Elefantes de columpiaban`)
}
