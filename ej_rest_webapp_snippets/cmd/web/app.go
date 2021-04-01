package web

import (
	"log"
	"pastein/pkg/models"
	"text/template"

	"github.com/golangcollege/sessions"
)

type contextKey string

var contextKeyUser = contextKey("user")

type Application struct {
	ErrorLog      *log.Logger
	InfoLog       *log.Logger
	Session       *sessions.Session
	Snippets      SnippetsInterfaces
	TemplateCache map[string]*template.Template
	Users         UsersInterfaces
}

type SnippetsInterfaces interface {
	Insert(*models.SnippetRequest) (int, error)
	Get(int) (*models.Snippet, error)
	Latest() ([]*models.Snippet, error)
}

type UsersInterfaces interface {
	Insert(*models.UserRequest) error
	Authenticate(*models.UserRequest) (int, error)
	Get(int) (*models.User, error)
}

var App = &Application{}
