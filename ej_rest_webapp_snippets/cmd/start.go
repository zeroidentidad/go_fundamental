package cmd

import (
	"flag"
	"pastein/cmd/web"
	"pastein/pkg/mysql"
	"time"

	"github.com/golangcollege/sessions"
)

func Start() {
	// default conf
	addr := flag.String("addr", ":4000", "HTTP network addr")
	dsn := flag.String("dsn", "webusr:passx123@tcp(localhost:3306)/snippetbox?parseTime=true", "MySQL db")
	secret := flag.String("secret", "z3Roh+pPbnzHbS*+9Pk8qGWhTzbpa@jf", "Secret session key")
	flag.Parse()

	db := web.Opendb(dsn)
	defer db.Close()

	templateCache, err := web.NewTemplateCache("./ui/html/")
	if err != nil {
		web.App.Logs().ErrorLog.Fatal(err)
	}

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour
	session.Secure = true
	// CSRF Protection opt
	// session.SameSite = http.SameSiteStrictMode

	app := &web.Application{
		ErrorLog:      web.App.Logs().ErrorLog,
		InfoLog:       web.App.Logs().InfoLog,
		Session:       session,
		Snippets:      &mysql.SnippetModel{DB: db},
		Users:         &mysql.UserModel{DB: db},
		TemplateCache: templateCache,
	}

	web.Serve(app.Routes(), addr)
}
