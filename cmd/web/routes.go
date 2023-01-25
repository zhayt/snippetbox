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

	// routes for user signup
	mux.Get("/user/signup", dynamicMiddelware.ThenFunc(app.singupUserForm))
	mux.Post("/user/signup", dynamicMiddelware.ThenFunc(app.singupUser))

	// user signin
	mux.Get("/user/login", dynamicMiddelware.ThenFunc(app.singinUserForm))
	mux.Post("/user/login", dynamicMiddelware.ThenFunc(app.singinUser))

	mux.Post("/user/logout", dynamicMiddelware.ThenFunc(app.logoutUser))

	// static files
	mux.Get("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static"))))

	return standardMiddleware.Then(mux)
}
