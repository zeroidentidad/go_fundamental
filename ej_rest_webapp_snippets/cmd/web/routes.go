package web

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

// *http.ServeMux. -> http.Handler
func (app *Application) Routes() http.Handler {
	stdMiddlewares := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dncMiddleware := alice.New(app.Session.Enable, noSurf, app.authenticate)

	mux := pat.New()
	mux.Get("/", dncMiddleware.ThenFunc(app.home))
	mux.Get("/snippet/create", dncMiddleware.Append(app.reqAuthUser).ThenFunc(app.createSnippetForm))
	mux.Post("/snippet/create", dncMiddleware.Append(app.reqAuthUser).ThenFunc(app.createSnippet))
	mux.Get("/snippet/:id", dncMiddleware.ThenFunc(app.showSnippet))
	mux.Get("/user/signup", dncMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dncMiddleware.ThenFunc(app.signupUser))
	mux.Get("/user/login", dncMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dncMiddleware.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dncMiddleware.Append(app.reqAuthUser).ThenFunc(app.logoutUser))
	mux.Get("/ping", http.HandlerFunc(ping))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return stdMiddlewares.Then(mux)
}
