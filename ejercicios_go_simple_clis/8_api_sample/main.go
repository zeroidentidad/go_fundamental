package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "query value: "+r.URL.Query().Get("tst"))
		//fmt.Fprintf(w, "http method: "+r.Method)
		//fmt.Fprintf(w, "http method: "+r.FormValue("testvalue"))
		fmt.Fprintf(w, "http method: "+r.Header.Get("testheader"))
	})

	http.ListenAndServe(":8080", nil)

}
