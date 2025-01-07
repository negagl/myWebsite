package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/negagl/myWebsite/internal/blog"
	"github.com/negagl/myWebsite/internal/project"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index.html")
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

	return r
}
