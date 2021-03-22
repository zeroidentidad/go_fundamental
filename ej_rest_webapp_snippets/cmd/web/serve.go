package web

import (
	"crypto/tls"
	"net/http"
	"time"
)

// *http.ServeMux. -> http.Handler
func serve(mux http.Handler, addr *string) {
	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     l.logs().errorLog,
		Handler:      mux,
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	l.logs().infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	l.logs().errorLog.Fatal(err)
}
