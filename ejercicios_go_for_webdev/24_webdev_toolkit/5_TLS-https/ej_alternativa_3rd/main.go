package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	http.HandleFunc("/", watahea)

	m := &autocert.Manager{
		Cache:      autocert.DirCache("secret-dir"),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("localhost", "127.0.0.1"),
	}
	if err := m.Cache; err != nil {
		log.Fatalln(err)
	}

	go http.ListenAndServe(":8080", nil)

	srv := &http.Server{
		Addr: ":10443",
		TLSConfig: &tls.Config{
			GetCertificate: m.GetCertificate,
		},
	}

	log.Fatalln(srv.ListenAndServeTLS("", ""))
}

func watahea(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hola TLS")
}
