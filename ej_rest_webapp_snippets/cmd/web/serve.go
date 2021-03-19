package web

import (
	"net/http"
)

// *http.ServeMux. -> http.Handler
func serve(mux http.Handler, addr *string) {
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: l.logs().errorLog,
		Handler:  mux,
	}

	l.logs().infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	l.logs().errorLog.Fatal(err)
}
