package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println(r.Header)

		mitoken := r.Header.Get("token")
		if len(mitoken) != 0 {
			fmt.Println(mitoken)
		}

		r.Header.Set("nombre", "valor")

		fmt.Println(r.Header)
		fmt.Fprintf(w, "Holas ...")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// test headers:
// curl http://localhost:8080/params -H "encabezadox:valorx"
// curl http://localhost:8080/params -H "token:xyz123"
