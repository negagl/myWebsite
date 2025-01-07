package blog

import (
	"encoding/json"
	"net/http"

	"github.com/negagl/myWebsite/helpers"
)

// Create a global instance of the blog store to use in the handlers
var store = NewStore()

func GetBlogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	store.mu.Lock()
	defer store.mu.Unlock()
	if err := json.NewEncoder(w).Encode(store.blogs); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func CreateBlog(w http.ResponseWriter, r *http.Request) {
	var newBlog Blog
	if err := ParseJSONToBlog(&newBlog, r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := ValidateBlog(newBlog); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	store.AddBlog(newBlog)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newBlog); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func GetBlogByID(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.ExtractIDFromPath(w, r, "blogs")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	blog, err := store.GetBlogByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusFound)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(blog); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func DeleteBlogByID(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.ExtractIDFromPath(w, r, "blogs")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := store.DeleteBlog(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(struct{}{}); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func UpdateBlogByID(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.ExtractIDFromPath(w, r, "blogs")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var updatedBody Blog
	if err := ParseJSONToBlog(&updatedBody, r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := ValidateBlog(updatedBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedBody.ID = id
	if err := store.UpdateBlog(updatedBody); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updatedBody); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
