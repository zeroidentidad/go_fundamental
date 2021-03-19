package web

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

// *http.ServeMux. -> http.Handler
func (app *application) routes() http.Handler {
	stdMiddlewares := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	mux := pat.New()

	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/snippet/create", http.HandlerFunc(app.createSnippetForm))
	mux.Post("/snippet/create", http.HandlerFunc(app.createSnippet))
	mux.Get("/snippet/:id", http.HandlerFunc(app.showSnippet))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return stdMiddlewares.Then(mux)
}
