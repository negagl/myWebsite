package blog

import (
	"errors"
)

// FindBlogByID finds a blog by its ID. Returns the blog and the index if found or -1 if not found.
func FindBlogByID(id int) (*Blog, int) {
	for i, blog := range blogs {
		if blog.ID == id {
			return &blog, i
		}
	}

	return nil, -1
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
