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

	// Restrict sub-tree path to match only /
	mux.HandleFunc("GET /{$}", app.home)

	// add wildcard pattern, id segment in the route
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)

	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	// Flow of control down the chain : recoverPanic -> logRequest -> commonHeaders -> servemux -> application handler
	// Flow of control back the chain : application handler -> servemux -> commonHeaders -> logRequest -> recoverPanic
	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
