package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/negagl/myWebsite/internal/blog"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "ok"}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	// Paths
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to my personal website!")
	})
	http.HandleFunc("/health", healthCheckHandler)
	http.HandleFunc("/blogs", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			blog.GetBlogs(w, r)
		case http.MethodPost:
			blog.CreateBlog(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/blogs/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			blog.GetBlogByID(w, r)
		case http.MethodDelete:
			blog.DeleteBlogByID(w, r)
		case http.MethodPut:
			blog.UpdateBlogByID(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Run server
	fmt.Println("Server is running on http://localhost:8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("There was an error starting the server")
		return
	}
}
