package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/testauth", func(w http.ResponseWriter, r *http.Request) {
		if authorize(r.Header.Get("Auth")) == "" {
			w.WriteHeader(401)
			return
		}
		fmt.Fprintf(w, "passed")
	})

	http.HandleFunc("/testmethod", func(w http.ResponseWriter, r *http.Request) {
		if authorize(r.Header.Get("Auth")) == "" {
			w.WriteHeader(401)
			return
		}
		if r.Method != "GET" {
			w.WriteHeader(405)
			return
		}
		fmt.Fprintf(w, "passed")
	})

	http.ListenAndServe(":8080", nil)
}

func authorize(token string) string {
	if token == "test" {
		return "userid"
	}
	return ""
}
