package cmd

import (
	"flag"
	"net/http"
)

func serve(mux *http.ServeMux) {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: l.logs().errorLog,
		Handler:  mux,
	}

	l.logs().infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	l.logs().errorLog.Fatal(err)
}
