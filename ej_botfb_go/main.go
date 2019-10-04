package main

import (
	"log"
	"net/http"
)

func main() {
	getConfig()

	http.HandleFunc("/", saluda)
	//log.Println(http.ListenAndServe("localhost:8080", nil))
	log.Printf("Serv iniciado https://localhost:%v", c.Port)
	err := http.ListenAndServeTLS(c.Port, c.CertPem, c.KeyPem, nil)
	if err != nil {
		log.Println(err)
	}
}

func saluda(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola x"))
}
