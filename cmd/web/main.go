package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Create a handler for serving static files at ./ui/static/ directory from the project root directory
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Register the fileServer to /static/ URL path, we must strip the /static from the URL
	// If we didn't strip it, the fileServer handler path will look for a file in ./ui/static//static/ which doesn't exist
	// Using the StripPrefix, the path that passed to fileServer handler will become ./ui/static// which is a valid path
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /", home)
	mux.HandleFunc("GET /snippet/view", snippetView)
	mux.HandleFunc("POST /snippet/create", snippetCreate)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
