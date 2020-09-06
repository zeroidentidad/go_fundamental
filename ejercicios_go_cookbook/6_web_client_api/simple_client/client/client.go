package client

import (
	"crypto/tls"
	"net/http"
)

// Setup configura nuestro cliente y
// redefine el DefaultClient global
func Setup(isSecure, nop bool) *http.Client {
	c := http.DefaultClient

	// A veces, para realizar pruebas, necesario
	// desactivar la verificación SSL
	if !isSecure {
		c.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
			},
		}
	}
	if nop {
		c.Transport = &NopTransport{}
	}
	http.DefaultClient = c
	return c
}

// NopTransport es un transporte sin operaciones (No-Op Transport)
type NopTransport struct {
}

// RoundTrip implementa la interfaz RoundTripper
func (n *NopTransport) RoundTrip(*http.Request) (*http.Response, error) {
	// notar que esta es una respuesta no inicializada
	// si se está buscando encabezados, etc.
	return &http.Response{StatusCode: http.StatusTeapot}, nil
}
