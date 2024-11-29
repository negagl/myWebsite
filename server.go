package main

import (
	"fmt"
	"net/http"
)

var blogs = []string{"Blog Post 1", "Blog Post 2", "Blog Post 3"}

func Server(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.WriteHeader(http.StatusOK)
			serveContent(w, HomePageContent())
		case "/blogs":
			w.WriteHeader(http.StatusOK)
			serveContent(w, BlogsContent())
		}
	case http.MethodPost:
		if r.URL.Path == "/blogs" {
			handleCreateBlog(w, r)
		}
	default:
		http.NotFound(w, r)
	}
}

func serveContent(w http.ResponseWriter, content string) {
	fmt.Fprint(w, content)
}

func handleCreateBlog(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	newBlog := string(body)
	blogs = append(blogs, newBlog)

	w.WriteHeader(http.StatusCreated)
}
