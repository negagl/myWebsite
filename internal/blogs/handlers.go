package blogs

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func GetBlogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogs)
}

func CreateBlog(w http.ResponseWriter, r *http.Request) {
	var newBlog Blog

	if err := json.NewDecoder(r.Body).Decode(&newBlog); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := ValidateBlog(newBlog); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	blog, _ := FindBlogByID(id)
	if blog == nil {
		// Blog not found
		http.Error(w, "Blog not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusFound)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blog)
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
	blog, i := FindBlogByID(id)
	if blog == nil {
		http.Error(w, "Blog not found", http.StatusNotFound)
		return
	}

	blogs = append(blogs[:i], blogs[i+1:]...)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blog)
}

func UpdateBlogByID(w http.ResponseWriter, r *http.Request) {
	// Get ID
	idString := r.URL.Path[len("/blogs/"):]
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid blog ID", http.StatusBadRequest)
		return
	}

	// Decode the body
	var updatedBody Blog
	if err := json.NewDecoder(r.Body).Decode(&updatedBody); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate blog
	if err := ValidateBlog(updatedBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update blog
	_, i := FindBlogByID(id)
	if i == -1 {
		http.Error(w, "Blog not found", http.StatusNotFound)
		return
	}

	blogs[i].Title = updatedBody.Title
	blogs[i].Content = updatedBody.Content

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogs[i])
}
