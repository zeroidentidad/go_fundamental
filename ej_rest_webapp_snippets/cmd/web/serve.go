package web

import (
	"net/http"
)

func serve(mux *http.ServeMux, addr *string) {
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: l.logs().errorLog,
		Handler:  mux,
	}

	l.logs().infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	l.logs().errorLog.Fatal(err)
}
