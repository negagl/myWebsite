package blog

import (
	"encoding/json"
	"errors"
	"net/http"
)

// ValidateBlog checks if blog data is valid.
func ValidateBlog(blog Blog) error {
	if blog.Title == "" {
		return errors.New("Title cannot be empty")
	}
	if blog.Content == "" {
		return errors.New("Content cannot be empty")
	}
	for _, existingBlog := range store.blogs {
		if existingBlog.ID == blog.ID {
			return errors.New("ID must be unique")
		}
	}
	return nil
}

// ParseJSONToBlog parses the request body to a Blog struct.
func ParseJSONToBlog(blogToParse *Blog, r *http.Request) error {
	if r.Body == nil {
		return errors.New("Request body cannot be empty")
	}
	if err := json.NewDecoder(r.Body).Decode(blogToParse); err != nil {
		return errors.New("Invalid JSON")
	}
	return nil
}
