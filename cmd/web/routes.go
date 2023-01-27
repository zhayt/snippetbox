package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable, noSurf, app.authenticate)

	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(http.HandlerFunc(app.home)))
	mux.Get("/snippet/create", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.createSnippetForm))
	mux.Post("/snippet/create", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.createSnippet))
	mux.Get("/snippet/:id", dynamicMiddleware.ThenFunc(app.showSnippet))

	// routes for user signup
	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.singupUserForm))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.singupUser))

	// user signin
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))

	mux.Post("/user/logout", dynamicMiddleware.ThenFunc(app.logoutUser))

	// Register the ping handler function as the handler for the GET /ping route.
	mux.Get("/ping", http.HandlerFunc(ping))

	// static files
	mux.Get("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static"))))

	return standardMiddleware.Then(mux)
}
