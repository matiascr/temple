package main

import (
	"html/template"
	"net/http"
)

type Model struct {
	Template string
}

const templatesPath = "templates/*"

var (
	model     Model
	templates *template.Template
	router    *http.ServeMux
)

func init() {
	// Parse all the templates
	templates, _ = template.ParseGlob(templatesPath)
	// Initialize the router
	router = http.NewServeMux()
	// Add static css path to FS
	router.Handle("/css/", http.FileServer(http.Dir(".")))
}

func main() {
	router.HandleFunc("/", index)

	http.ListenAndServe(":8080", router)
}

func index(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
}
