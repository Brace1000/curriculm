package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var tmpl *template.Template

// Initialize templates on server startup
func init() {
	var err error
	tmpl, err = template.ParseFiles(filepath.Join("templates", "index.html"))
	if err != nil {
		log.Fatalf("failed to parse template: %v", err)
	}
	
}

// HomeHandler serves the index.html file
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure only the root path is served by this handler
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		log.Printf("template execution error: %v", err)
	}
}
