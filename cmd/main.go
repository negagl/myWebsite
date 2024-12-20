package main

import (
	"encoding/json"
	"fmt"
	"github.com/negagl/myWebsite/internal/blogs"
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "ok"}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	jsonResponse, _ := json.Marshal(response)
	w.Write(jsonResponse)
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
			blogs.GetBlogs(w, r)
		case http.MethodPost:
			blogs.CreateBlog(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/blogs/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			blogs.GetBlogByID(w, r)
		case http.MethodDelete:
			blogs.DeleteBlogByID(w, r)
		case http.MethodPut:
			blogs.UpdateBlogByID(w, r)
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
