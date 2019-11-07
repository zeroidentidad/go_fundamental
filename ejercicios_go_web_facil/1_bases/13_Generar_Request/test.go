package main

import (
	"fmt"
	"log"
	"net/http"
)

//GET, POST, PUT, DELETE

func main() {
	http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request) {

		//fmt.Println(r.URL.RawQuery)

		fmt.Println(r.URL)
		valores := r.URL.Query()

		valores.Del("otro") // quitar parametro innecesario
		valores.Add("personal", "mivalor")
		valores.Add("personal2", "mivalor2")
		valores.Add("personal3", "mivalor3")

		r.URL.RawQuery = valores.Encode()

		//fmt.Fprint(w, r.URL)
		fmt.Fprintf(w, r.URL.String())
		fmt.Println(r.URL)

	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// ejecutar 1ro
