package blog

import (
	"encoding/json"
	"errors"
	"net/http"
)

// FindBlogByID finds a blog by its ID. Returns the blog and the index if found or -1 if not found.
func FindBlogByID(id int) (*Blog, int, error) {
	for i, blog := range blogs {
		if blog.ID == id {
			return &blog, i, nil
		}
	}

	return nil, -1, errors.New("Blog not found")
}

// ValidateBlog checks if a blog data is valid.
func ValidateBlog(blog Blog) error {
	if blog.Title == "" {
		return errors.New("Title cannot be empty")
	}

	if blog.Content == "" {
		return errors.New("Content cannot be empty")
	}

	for _, existingBlog := range blogs {
		if existingBlog.ID == blog.ID {
			return errors.New("ID must be unique")
		}
	}

	return nil
}

// ParseJSONToBlog parses the body of a request to a Blog struct
func ParseJSONToBlog(blogToParse *Blog, r *http.Request) error {
	if r.Body == nil {
		return errors.New("Request body cannot be empty")
	}

	if err := json.NewDecoder(r.Body).Decode(&blogToParse); err != nil {
		return errors.New("Invalid JSON")
	}

	return nil
}
