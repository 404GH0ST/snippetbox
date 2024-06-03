package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// Create a handler for serving static files at ./ui/static directory from the project root directory
	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static")})

	// Register the fileServer to /static/ URL path, we must strip the /static from the URL
	// If we didn't strip it, the fileServer handler path will look for a file in ./ui/static/static/ which doesn't exist
	// Using the StripPrefix, the path that passed to fileServer handler will become ./ui/static/ which is a valid path
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /", app.home)
	mux.HandleFunc("GET /snippet/view", app.snippetView)
	mux.HandleFunc("POST /snippet/create", app.snippetCreate)

	return mux
}
