package main

import (
	"fmt"
	"log"
	"net/http"
)

//GET, POST, PUT, DELETE

func main() {
	http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request) {

		//fmt.Println(r.URL)
		//fmt.Println(r.URL.RawQuery)
		fmt.Println(r.URL.Query()) //Map clave - valor

		name := r.URL.Query().Get("name")
		if len(name) != 0 {
			fmt.Println(name)
		}

		parametro := r.URL.Query().Get("parametro")
		if len(parametro) != 0 {
			fmt.Println(parametro)
		}

	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// test metodos:
// curl http://localhost:8080/params?name=Jesus -X GET
// curl http://localhost:8080/params?name=Jesus341 -X GET
// curl "http://localhost:8080/params?name=Jesus&parametro=valor123" -X GET
// curl http://localhost:8080/params -X POST -i
