package main

import (
	"log"
	"net/http"
)

func main() {
	/*http.Handle("/", http.FileServer(http.Dir("public")))

	log.Println("Server ejecutandose en :8080")

	log.Println(http.ListenAndServe(":8080", nil))*/

	multiplexor := http.NewServeMux()
	filesystem := http.FileServer(http.Dir("public"))
	multiplexor.Handle("/", filesystem)
	log.Println("Server ejecutandose en :8080")
	log.Println(http.ListenAndServe(":8080", multiplexor))
}
