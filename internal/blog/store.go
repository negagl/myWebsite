package blog

import (
	"errors"
	"sync"
)

// Store maintains an in-memory list of blogs.
// We'll use an in-memory solution for now and switch to SQLite later.
type Store struct {
	blogs []Blog
	mu    sync.Mutex
}

// NewStore creates a new instance of Store.
func NewStore() *Store {
	return &Store{
		blogs: blogs, // Initialize with the test list
	}
}

// AddBlog adds a new blog.
func (s *Store) AddBlog(b Blog) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.blogs = append(s.blogs, b)
}

// GetBlogByID gets a blog by its ID.
func (s *Store) GetBlogByID(id int) (Blog, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, b := range s.blogs {
		if b.ID == id {
			return b, nil
		}
	}
	return Blog{}, errors.New("blog not found")
}

// UpdateBlog updates an existing blog.
func (s *Store) UpdateBlog(updated Blog) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, b := range s.blogs {
		if b.ID == updated.ID {
			s.blogs[i] = updated
			return nil
		}
	}
	return errors.New("blog not found")
}

// DeleteBlog deletes a blog by its ID.
func (s *Store) DeleteBlog(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, b := range s.blogs {
		if b.ID == id {
			s.blogs = append(s.blogs[:i], s.blogs[i+1:]...)
			return nil
		}
	}
	return errors.New("blog not found")
}
