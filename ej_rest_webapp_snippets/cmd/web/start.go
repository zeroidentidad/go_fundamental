package cmd

import (
	"log"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

var l = &application{}

func Start() {
	app := &application{
		errorLog: l.logs().errorLog,
		infoLog:  l.logs().infoLog,
	}

	serve(app.routes())
}
