package web

import (
	"flag"
	"log"
	"pastein/pkg/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *mysql.SnippetModel
}

var l = &application{}

func Start() {
	addr := flag.String("addr", ":4000", "HTTP network addr")
	dsn := flag.String("dsn", "webusr:passx123@tcp(localhost:3306)/snippetbox?parseTime=true", "MySQL db")
	flag.Parse()

	db := opendb(dsn)
	defer db.Close()

	app := &application{
		errorLog: l.logs().errorLog,
		infoLog:  l.logs().infoLog,
		snippets: &mysql.SnippetModel{DB: db},
	}

	serve(app.routes(), addr)
}
