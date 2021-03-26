package web

import (
	"log"
	"os"
)

func (app *Application) Logs() (logs *Application) {
	/*f, err := os.OpenFile("/tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()*/

	//f -> os.Stdout
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	logs = &Application{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
	}

	return logs
}
