package web

import (
	"log"
	"pastein/pkg/mysql"
	"text/template"

	"github.com/golangcollege/sessions"
)

type contextKey string

var contextKeyUser = contextKey("user")

type Application struct {
	ErrorLog      *log.Logger
	InfoLog       *log.Logger
	Session       *sessions.Session
	Snippets      *mysql.SnippetModel
	TemplateCache map[string]*template.Template
	Users         *mysql.UserModel
}

var App = &Application{}
