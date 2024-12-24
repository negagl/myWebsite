package blogs

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
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

func generateValidTokenForTest() string {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Username: "admin",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtSecret)
	return tokenString
}
