package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Make sure home handler can only be accessed from / path because of the subtree pattern that acts like catch-all
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	// Define required template file
	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}

	// Parse all the templates file, using variadic argument so that it reads all the slice element
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Render the base template, because in base template it will invoke other template
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	// Convert GET id parameter to an int
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	// Make sure it's a valid positive integer
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Displaying specific snippet ID : %d...", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Snippet created successfully"))
}
