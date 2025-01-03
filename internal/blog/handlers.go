package blog

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/negagl/myWebsite/helpers"
)

func GetBlogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(blogs); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func CreateBlog(w http.ResponseWriter, r *http.Request) {
	// parse blog
	var newBlog Blog
	if err := ParseJSONToBlog(&newBlog, r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// validations
	if err := ValidateBlog(newBlog); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Add blog
	blogs = append(blogs, newBlog)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newBlog); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func GetBlogByID(w http.ResponseWriter, r *http.Request) {
	// Extract ID from the URL
	id, err := helpers.ExtractIDFromPath(w, r, "blogs")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Search for the blog
	blog, _, err := FindBlogByID(id)
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
	// Extract ID from URL
	id, err := helpers.ExtractIDFromPath(w, r, "blogs")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Search for the blog
	blog, i, err := FindBlogByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	blogs = append(blogs[:i], blogs[i+1:]...)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(blog); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func UpdateBlogByID(w http.ResponseWriter, r *http.Request) {
	// Get ID
	id, err := helpers.ExtractIDFromPath(w, r, "blogs")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// parse the body
	var updatedBody Blog
	if err := ParseJSONToBlog(&updatedBody, r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validations
	if err := ValidateBlog(updatedBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update blog
	_, i, err := FindBlogByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	blogs[i].Title = updatedBody.Title
	blogs[i].Content = updatedBody.Content

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(blogs[i]); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if creds.Username != "admin" || creds.Password != "secret" {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Create JWT
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write([]byte(`{"token":""` + tokenString + `"}`)); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
