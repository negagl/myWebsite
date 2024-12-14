package blogs

import (
	"encoding/json"
	"net/http"
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
	jsonResponse, _ := json.Marshal(blogs)
	w.Write(jsonResponse)
}

func CreateBlog(w http.ResponseWriter, r *http.Request) {
	var newBlog Blog
	if err := json.NewDecoder(r.Body).Decode(&newBlog); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Add blog
	blogs = append(blogs, newBlog)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	jsonResponse, _ := json.Marshal(newBlog)
	w.Write(jsonResponse)
}
