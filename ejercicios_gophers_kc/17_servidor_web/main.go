package main

import "net/http"
import "fmt"

func handlerHola(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v, %v\n", r.Method, r.Host)
	fmt.Fprintln(w, "Hola Mundo")
}

func main() {
	http.HandleFunc("/hola", func(w http.ResponseWriter, r *http.Request) {
		go handlerHola(w, r)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error creando el servidor")
		fmt.Println(err)
	}
}
