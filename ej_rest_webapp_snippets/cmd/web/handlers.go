package web

import (
	"fmt"
	"net/http"
	"pastein/pkg/forms"
	"pastein/pkg/models"
	"strconv"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	s, err := app.Snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{Snippets: s})
}

func (app *Application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.Snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "show.page.tmpl", &templateData{Snippet: s})
}

func (app *Application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", &templateData{Form: forms.New(nil)})
}

func (app *Application) createSnippet(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLength("title", 100)
	form.PermittedValues("expires", "365", "7", "1")

	if !form.Valid() {
		app.render(w, r, "create.page.tmpl", &templateData{Form: form})
		return
	}

	request := &models.SnippetRequest{
		Title:   form.Get("title"),
		Content: form.Get("content"),
		Expires: form.Get("expires"),
	}

	id, err := app.Snippets.Insert(request)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.Session.Put(r, "flash", "Snippet creado correctamente!")

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}

func (app *Application) signupUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "signup.page.tmpl", &templateData{Form: forms.New(nil)})
}

func (app *Application) signupUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("name", "email", "password")
	form.MatchesPattern("email", forms.EmailRX)
	form.MinLength("password", 8)

	if !form.Valid() {
		app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
		return
	}

	request := &models.UserRequest{
		Name:     form.Get("name"),
		Email:    form.Get("email"),
		Password: form.Get("password"),
	}

	err = app.Users.Insert(request)
	if err == models.ErrDuplicateEmail {
		form.Errors.Add("email", "Email ya está en uso")
		app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	app.Session.Put(r, "flash", "Registro exitoso. Puede iniciar sesión.")

	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (app *Application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.tmpl", &templateData{Form: forms.New(nil)})
}

func (app *Application) loginUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	request := &models.UserRequest{
		Email:    form.Get("email"),
		Password: form.Get("password"),
	}

	id, err := app.Users.Authenticate(request)
	if err == models.ErrInvalidCredentials {
		form.Errors.Add("generic", "Datos de acceso incorrectos")
		app.render(w, r, "login.page.tmpl", &templateData{Form: form})
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	app.Session.Put(r, "userID", id)

	http.Redirect(w, r, "/snippet/create", http.StatusSeeOther)
}

func (app *Application) logoutUser(w http.ResponseWriter, r *http.Request) {
	app.Session.Remove(r, "userID")
	app.Session.Put(r, "flash", "Has cerrado sesión!")

	http.Redirect(w, r, "/", 303)
}
