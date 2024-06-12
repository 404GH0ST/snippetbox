package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// Create a handler for serving static files at ./ui/static directory from the project root directory
	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static")})

	// Register the fileServer to /static/ URL path, we must strip the /static from the URL
	// If we didn't strip it, the fileServer handler path will look for a file in ./ui/static/static/ which doesn't exist
	// Using the StripPrefix, the path that passed to fileServer handler will become ./ui/static/ which is a valid path
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	// Because the alice.ThenFunc() methods returns a http.handler (rather than http.HandlerFunc) we also
	// need to switch to registering the route using the mux.handle() method.
	// Restrict sub-tree path to match only /
	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home))

	// add wildcard pattern, id segment in the route
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.snippetView))
	mux.Handle("GET /user/signup", dynamic.ThenFunc(app.userSignup))
	mux.Handle("POST /user/signup", dynamic.ThenFunc(app.userSignupPost))
	mux.Handle("GET /user/login", dynamic.ThenFunc(app.userLogin))
	mux.Handle("POST /user/login", dynamic.ThenFunc(app.userLoginPost))

	protected := dynamic.Append(app.requireAuthentication)

	mux.Handle("GET /snippet/create", protected.ThenFunc(app.snippetCreate))
	mux.Handle("POST /snippet/create", protected.ThenFunc(app.snippetCreatePost))
	mux.Handle("POST /user/logout", protected.ThenFunc(app.userLogoutPost))

	// Flow of control down the chain : recoverPanic -> logRequest -> commonHeaders -> servemux -> application handler
	// Flow of control back the chain : application handler -> servemux -> commonHeaders -> logRequest -> recoverPanic
	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
