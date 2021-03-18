package web

import (
	"flag"
	"log"
	"pastein/pkg/mysql"
	"text/template"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	snippets      *mysql.SnippetModel
	templateCache map[string]*template.Template
}

var l = &application{}

func Start() {
	addr := flag.String("addr", ":4000", "HTTP network addr")
	dsn := flag.String("dsn", "webusr:passx123@tcp(localhost:3306)/snippetbox?parseTime=true", "MySQL db")
	flag.Parse()

	db := opendb(dsn)
	defer db.Close()

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		l.logs().errorLog.Fatal(err)
	}

	app := &application{
		errorLog:      l.logs().errorLog,
		infoLog:       l.logs().infoLog,
		snippets:      &mysql.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	serve(app.routes(), addr)
}
