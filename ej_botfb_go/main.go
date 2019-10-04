package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", saluda)
	//log.Println(http.ListenAndServe("localhost:8080", nil))
	log.Println("Serv iniciado https://...")
	err := http.ListenAndServeTLS(":443", "./certificados/cert.pem", "./certificados/key.pem", nil)
	if err != nil {
		log.Println(err)
	}
}

func saluda(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola x"))
}
