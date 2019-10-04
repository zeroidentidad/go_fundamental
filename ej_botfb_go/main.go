package main

import (
	"log"
	"net/http"
)

func main() {
	getConfig()

	http.HandleFunc("/", saluda)
	http.HandleFunc("/webhook", webhook)

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

func webhook(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		vertoken := r.URL.Query().Get("hub.verify_token")
		if vertoken == c.MyToken {
			hubc := r.URL.Query().Get("hub.challenge")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(hubc))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error en la verificación"))
		return
	}
}
