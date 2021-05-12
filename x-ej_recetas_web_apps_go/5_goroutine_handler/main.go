package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("hola, gophers")
		}()

		return

	})

	log.Println("Running on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
