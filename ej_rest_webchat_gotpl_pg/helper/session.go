package helper

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"webchat/db"
)

// para redirigir a la página del mensaje de error
func ErrorMessage(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

// comprobar si el usuario está conectado y tiene una sesión, si no, err
func Session(writer http.ResponseWriter, request *http.Request) (sess db.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = db.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}

	return
}

// analizar plantillas HTML, pasar lista de nombres de archivo y obtener plantilla
func ParseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))

	return
}

func GenerateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	_ = templates.ExecuteTemplate(writer, "layout", data)
}
