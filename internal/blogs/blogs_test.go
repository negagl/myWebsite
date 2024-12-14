package blogs

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetBlogs(t *testing.T) {
	// Request
	req, err := http.NewRequest(http.MethodGet, "/blogs", nil)
	if err != nil {
		t.Fatalf("Could not create the request: %v", err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBlogs)

	// Execution
	handler.ServeHTTP(rec, req)

	// Assertions
	if rec.Code != http.StatusOK {
		t.Errorf("Unexpected status code: got %d, want %d", rec.Code, http.StatusOK)
	}

	expected := `[{"id":1,"title":"First Blog","content":"This is the first blog post"}]`
	if rec.Body.String() != expected {
		t.Errorf("Unexpected response body: got %s, want %s", rec.Body.String(), expected)
	}
}

func TestCreateBlog(t *testing.T) {
	body := `{"id":2,"title":"Second Blog","content":"This is the second blog post"}`
	req, err := http.NewRequest(http.MethodPost, "/blogs", strings.NewReader(body))
	if err != nil {
		t.Fatalf("Could not create the request: %v", err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateBlog)

	handler.ServeHTTP(rec, req)

	// Assertions
	if rec.Code != http.StatusCreated {
		t.Errorf("Unexpected status code: got %d, want %d", rec.Code, http.StatusCreated)
	}

	expected := `{"id":2,"title":"Second Blog","content":"This is the second blog post"}`
	if rec.Body.String() != expected {
		t.Errorf("Unexpected body response: got %s, want %s", rec.Body.String(), expected)
	}
}
