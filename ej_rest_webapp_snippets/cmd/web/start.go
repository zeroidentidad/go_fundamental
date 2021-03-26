package web

import (
	"flag"
	"log"
	"pastein/pkg/mysql"
	"text/template"
	"time"

	"github.com/golangcollege/sessions"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	session       *sessions.Session
	snippets      *mysql.SnippetModel
	templateCache map[string]*template.Template
	users         *mysql.UserModel
}

var l = &application{}

func Start() {
	addr := flag.String("addr", ":4000", "HTTP network addr")
	dsn := flag.String("dsn", "webusr:passx123@tcp(localhost:3306)/snippetbox?parseTime=true", "MySQL db")
	secret := flag.String("secret", "z3Roh+pPbnzHbS*+9Pk8qGWhTzbpa@jf", "Secret session key")
	flag.Parse()

	db := opendb(dsn)
	defer db.Close()

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		l.logs().errorLog.Fatal(err)
	}

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour
	session.Secure = true
	// CSRF Protection opt
	// session.SameSite = http.SameSiteStrictMode

	app := &application{
		errorLog:      l.logs().errorLog,
		infoLog:       l.logs().infoLog,
		session:       session,
		snippets:      &mysql.SnippetModel{DB: db},
		users:         &mysql.UserModel{DB: db},
		templateCache: templateCache,
	}

	serve(app.routes(), addr)
}
