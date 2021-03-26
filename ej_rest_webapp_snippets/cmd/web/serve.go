package web

import (
	"crypto/tls"
	"net/http"
	"time"
)

// *http.ServeMux. -> http.Handler
func Serve(mux http.Handler, addr *string) {
	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     App.Logs().ErrorLog,
		Handler:      mux,
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	App.Logs().InfoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	App.Logs().ErrorLog.Fatal(err)
}
