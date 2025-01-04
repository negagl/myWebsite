package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/negagl/myWebsite/internal/blog"
	"github.com/negagl/myWebsite/internal/project"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "ok"}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to my personal website!")
	})

	r.Get("/health", healthCheckHandler)

	r.Route("/blogs", func(r chi.Router) {
		r.Get("/", blog.GetBlogs)
		r.Get("/{id}", blog.GetBlogByID)
		r.Post("/", blog.CreateBlog)
		r.Put("/{id}", blog.UpdateBlogByID)
		r.Delete("/{id}", blog.DeleteBlogByID)
	})

	r.Route("/projects", func(r chi.Router) {
		r.Get("/", project.GetProjects)
		r.Get("/{id}", project.GetProjectByID)
		r.Post("/", project.CreateProject)
		r.Put("/{id}", project.UpdateProject)
		r.Delete("/{id}", project.DeleteProject)
	})

	// Run server
	fmt.Println("Server is running on http://localhost:8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("There was an error starting the server")
		return
	}
}
