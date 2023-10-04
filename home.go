package main

import (
	_ "embed"
	"net/http"
)

//go:embed home.html
var homeHTML []byte

func (a *Amadeus) HomeHandler(w http.ResponseWriter, r *http.Request) {
	// The "/" pattern matches everything, so we need to check
	// that we're at the root here.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Set the content type to HTML
	w.Header().Set("Content-Type", "text/html")
	defer r.Body.Close()

	// Write the embedded HTML file content to the response writer
	_, err := w.Write(homeHTML)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
