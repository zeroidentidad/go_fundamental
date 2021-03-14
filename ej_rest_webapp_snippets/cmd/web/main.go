package cmd

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func Start() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	/*f, ferr := os.OpenFile("/tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	if ferr != nil {
		log.Fatal(ferr)
	}
	defer f.Close()*/

	//f -> os.Stdout
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
