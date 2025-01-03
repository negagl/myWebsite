package blog

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	tests := []struct {
		name           string
		body           string
		expectedStatus int
		expectedError  string
	}{
		{
			name:           "Valid Credentials",
			body:           `{"username":"admin","password":"secret"}`,
			expectedStatus: http.StatusOK,
			expectedError:  "",
		},
		{
			name:           "Invalid Credentials",
			body:           `{"username":"wrong","password":"wrong"}`,
			expectedStatus: http.StatusUnauthorized,
			expectedError:  "Invalid username or password",
		},
		{
			name:           "Empty Body",
			body:           "{}",
			expectedStatus: http.StatusUnauthorized,
			expectedError:  "Invalid username or password",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/login", strings.NewReader(tc.body))
			if err != nil {
				t.Fatalf("Could not create the request")
			}

			rec := httptest.NewRecorder()
			Login(rec, req)

			if rec.Code != tc.expectedStatus {
				t.Errorf("Unexpected status code: got %d, want %d", rec.Code, tc.expectedStatus)
			}

			if tc.expectedError != "" && rec.Body.String() != tc.expectedError+"\n" {
				t.Errorf("unexpected error: got %q, want %q", rec.Body.String(), tc.expectedError)
			}
		})
	}
}

func TestGetBlogs(t *testing.T) {
	// Request
	req, err := http.NewRequest(http.MethodGet, "/blogs", nil)
	if err != nil {
		t.Fatalf("Could not create the request: %v", err)
	}

	rec := httptest.NewRecorder()
	GetBlogs(rec, req)

	// Assertions
	if rec.Code != http.StatusOK {
		t.Errorf("Unexpected status code: got %d, want %d", rec.Code, http.StatusOK)
	}

	expected := `[{"id":1,"title":"First Blog","content":"This is the first blog post"}]`
	if rec.Body.String() != expected+"\n" {
		t.Errorf("Unexpected response body: got %s, want %s", rec.Body.String(), expected)
	}
}

func TestCreateBlog(t *testing.T) {
	tests := []struct {
		name           string
		body           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Created correctly",
			body:           `{"id":2,"title":"Second Blog","content":"This is the second blog post"}`,
			expectedStatus: http.StatusCreated,
			expectedBody:   `{"id":2,"title":"Second Blog","content":"This is the second blog post"}`,
		},
		{
			name:           "Empty Title",
			body:           `{"id":3,"title":"","content":"Content"}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Title cannot be empty",
		},
		{
			name:           "Empty Content",
			body:           `{"id":3,"title":"Title","content":""}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Content cannot be empty",
		},
		{
			name:           "Duplicated ID",
			body:           `{"id":1,"title":"Title","content":"Content"}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "ID must be unique",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/blogs", strings.NewReader(tc.body))
			if err != nil {
				t.Fatalf("Could not create the request: %v", err)
			}

			rec := httptest.NewRecorder()
			CreateBlog(rec, req)

			if rec.Code != tc.expectedStatus {
				t.Errorf("Unexpected status code: got %d, want %d", rec.Code, tc.expectedStatus)
			}

			if rec.Body.String() != tc.expectedBody+"\n" {
				t.Errorf("Unexpected body response: got %q, want %q", rec.Body.String(), tc.expectedBody)
			}
		})
	}
}

func TestGetBlogByID(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Blog found",
			id:             "1",
			expectedStatus: http.StatusFound,
			expectedBody:   `{"id":1,"title":"First Blog","content":"This is the first blog post"}`,
		},
		{
			name:           "Blog not found",
			id:             "99",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Blog not found",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/blogs/"+tc.id, nil)
			if err != nil {
				t.Fatalf("Could not create the request: %v", err)
			}

			rec := httptest.NewRecorder()
			GetBlogByID(rec, req)

			if rec.Code != tc.expectedStatus {
				t.Errorf("Unexpected status code: got %d, want %d", rec.Code, tc.expectedStatus)
			}

			if rec.Body.String() != tc.expectedBody+"\n" {
				t.Errorf("Unexpected body response: got %q, want %q", rec.Body.String(), tc.expectedBody)
			}
		})
	}
}

func TestDeleteBlogByID(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Blog Found",
			id:             "1",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":1,"title":"First Blog","content":"This is the first blog post"}`,
		},
		{
			name:           "Blog not found",
			id:             "99",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Blog not found",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodDelete, "/blogs/"+tc.id, nil)
			if err != nil {
				t.Fatalf("Could not create the request: %v", err)
			}

			rec := httptest.NewRecorder()
			DeleteBlogByID(rec, req)

			if rec.Code != tc.expectedStatus {
				t.Errorf("Unexpected status code: got %d want %d", rec.Code, tc.expectedStatus)
			}

			if rec.Body.String() != tc.expectedBody+"\n" {
				t.Errorf("Unexpected response body: got %q, want %q", rec.Body.String(), tc.expectedBody)
			}
		})
	}
}

func TestUpdateBlogByID(t *testing.T) {
	blogs = []Blog{
		{ID: 1, Title: "First Blog", Content: "This is the first blog post"},
	}

	tests := []struct {
		name             string
		id               string
		body             string
		expectedStatus   int
		expectedResponse string
	}{
		{
			name:             "Blog found",
			id:               "1",
			body:             `{"title":"Third Blog","content":"This is the third blog"}`,
			expectedStatus:   http.StatusOK,
			expectedResponse: `{"id":1,"title":"Third Blog","content":"This is the third blog"}`,
		},
		{
			name:             "Blog not found",
			id:               "99",
			body:             `{"title":"Third Blog","content":"This is the third blog"}`,
			expectedStatus:   http.StatusNotFound,
			expectedResponse: "Blog not found",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPut, "/blogs/"+tc.id, strings.NewReader(tc.body))
			if err != nil {
				t.Fatalf("Could not create the request: %v", err)
			}

			rec := httptest.NewRecorder()
			UpdateBlogByID(rec, req)

			if rec.Code != tc.expectedStatus {
				t.Errorf("Unexpected status code: got %d, want %d", rec.Code, tc.expectedStatus)
			}

			if rec.Body.String() != tc.expectedResponse+"\n" {
				t.Errorf("Unexpected body response: got %q, want %q", rec.Body.String(), tc.expectedResponse)
			}
		})
	}
}
