package cmd

import (
	"flag"
	"log"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

var l = &application{}

func Start() {
	addr := flag.String("addr", ":4000", "HTTP network addr")
	dsn := flag.String("dsn", "webusr:passx123@tcp(localhost:3306)/snippetbox?parseTime=true", "MySQL db")
	flag.Parse()

	app := &application{
		errorLog: l.logs().errorLog,
		infoLog:  l.logs().infoLog,
	}

	db := opendb(dsn)
	defer db.Close()

	serve(app.routes(), addr)
}
