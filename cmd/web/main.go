package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	// Create new loggers, use bitwire OR operator | to combine flags
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	// Initializing a new instance of application struct, containing the dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}
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

	srv := http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", *addr) // Information message
	err := srv.ListenAndServe()

	// Fatal and Panic are recommended to use inside main function only
	errorLog.Fatal(err) // Error message, exit when occured
}
