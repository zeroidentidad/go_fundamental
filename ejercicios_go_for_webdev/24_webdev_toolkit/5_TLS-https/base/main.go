package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", watahea)
	http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
}

func watahea(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Este es un servidor TLS https de ejemplo (sin firmar).\n"))
}

// ir a https://localhost:10443/ o https://127.0.0.1:10443/

// lista puertos TCP:
// https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers

// Generar certificado sin firmar
// go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=somedomainname.com

// por ejemplo:
// go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=localhost

// WINDOWS
// windows puede tener problemas con go env GOROOT
// go run %(go env GOROOT)%/src/crypto/tls/generate_cert.go --host=localhost

// en vez de go env GOROOT
// usar la ruta al GO SDK donde sea que esté
