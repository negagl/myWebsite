package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

var templates = template.Must(template.New("").ParseFiles(
	"web/templates/base.html",
	"web/templates/components/header.html",
	"web/templates/components/footer.html",
	"web/templates/index.html",
	"web/templates/blogs.html",
))

func init() {
	// Show loaded templates
	for _, tmpl := range templates.Templates() {
		fmt.Println("Template loaded:", tmpl.Name())
	}
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "ok"}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(w, "base.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// Config reouter
	r := NewRouter()

	// Run server
	fmt.Println("Server is running on http://localhost:8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("There was an error starting the server:", err)
	}
}
