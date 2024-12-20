package blogs

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Blog struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var blogs = []Blog{
	{1, "First Blog", "This is the first blog post"},
}

func GetBlogs(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogs)
}

func CreateBlog(w http.ResponseWriter, r *http.Request) {
	var newBlog Blog

	if err := json.NewDecoder(r.Body).Decode(&newBlog); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate ID
	for _, blog := range blogs {
		if newBlog.ID == blog.ID {
			http.Error(w, "ID must be unique", http.StatusBadRequest)
			return
		}
	}

	// Validate Title
	if newBlog.Title == "" {
		http.Error(w, "Title cannot be empty", http.StatusBadRequest)
		return
	}

	// Validate Content
	if newBlog.Content == "" {
		http.Error(w, "Content cannot be empty", http.StatusBadRequest)
		return
	}

	// Add blog
	blogs = append(blogs, newBlog)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBlog)
}

func GetBlogByID(w http.ResponseWriter, r *http.Request) {
	// Extract ID from the URL
	idStr := r.URL.Path[len("/blogs/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid blog ID", http.StatusBadRequest)
		return
	}

	// Search for the blog
	for _, blog := range blogs {
		if blog.ID == id {
			w.WriteHeader(http.StatusFound)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(blog)
			return
		}
	}

	// Blog not found
	http.Error(w, "Blog not found", http.StatusNotFound)
}

func DeleteBlogByID(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL
	idStr := r.URL.Path[len("/blogs/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid blog ID", http.StatusBadRequest)
		return
	}

	// Search for the blog
	for i, blog := range blogs {
		if blog.ID == id {
			removedBlog := blogs[i]
			blogs = append(blogs[:i], blogs[i+1:]...)

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(removedBlog)
			return
		}
	}

	http.Error(w, "Blog not found", http.StatusNotFound)
}
