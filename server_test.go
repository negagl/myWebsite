package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomePage(t *testing.T) {
	t.Run("returns a basic HTML response on /", func(t *testing.T) {
		// Create a new http request to "/"
		request, _ := http.NewRequest(http.MethodGet, "/", nil)

		// Simulate response
		response := httptest.NewRecorder()

		// Call the server
		Server(response, request)

		//	get the response as text
		got := response.Body.String()
		want := "<h1>welcome to my site</h1>"

		//	check if all good
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestBlogPage(t *testing.T) {
	t.Run("returns a list of blogs", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/blogs", nil)
		response := httptest.NewRecorder()

		Server(response, request)

		got := response.Body.String()
		want := `
		<h1>My Blog</h1>
		<ul>
			<li>Blog Post 1</li><li>Blog Post 2</li><li>Blog Post 3</li>
		</ul>
		`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestBlogsCRUD(t *testing.T) {
	t.Run("create a blog post", func(t *testing.T) {
		// Create blog
		newPost := "Blog Post 4"
		request, _ := http.NewRequest(http.MethodPost, "/blogs", strings.NewReader(newPost))
		response := httptest.NewRecorder()

		Server(response, request)

		gotStatus := response.Code
		wantStatus := http.StatusCreated

		//Check if status is 201(created)
		if gotStatus != wantStatus {
			t.Errorf("got status %d, want status %d", gotStatus, wantStatus)
		}

		// Check if blog was actually created
		getRequest, _ := http.NewRequest(http.MethodGet, "/blogs", nil)
		getResponse := httptest.NewRecorder()

		Server(getResponse, getRequest)

		gotBody := getResponse.Body.String()
		if !strings.Contains(gotBody, fmt.Sprintf("<li>%s</li>", newPost)) {
			t.Errorf("Expected to find %q in the response, got %s", newPost, gotBody)
		}
	})
}
