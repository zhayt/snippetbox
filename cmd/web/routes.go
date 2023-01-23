package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddelware := alice.New(app.session.Enable)

	mux := pat.New()
	mux.Get("/", dynamicMiddelware.ThenFunc(http.HandlerFunc(app.home)))
	mux.Get("/snippet/create", dynamicMiddelware.ThenFunc(http.HandlerFunc(app.createSnippetForm)))
	mux.Post("/snippet/create", dynamicMiddelware.ThenFunc(http.HandlerFunc(app.createSnippet)))
	mux.Get("/snippet/:id", dynamicMiddelware.ThenFunc(http.HandlerFunc(app.showSnippet)))

	mux.Get("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static"))))

	return standardMiddleware.Then(mux)
}
