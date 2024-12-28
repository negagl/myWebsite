package projects

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProjects(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/projects", nil)
	if err != nil {
		t.Fatalf("Could not create the request: %v", err)
	}

	rec := httptest.NewRecorder()

	GetProjects(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Unexpected status code: got %d, want %d", rec.Code, http.StatusOK)
	}

	expectedResponse := `[{"id":1,"title":"Project 1","description":"This is the first project","url":"https://github.com/project1","status":"started"}]`
	if rec.Body.String() != expectedResponse+"\n" {
		t.Errorf("Unexpected response body: got %q, want %q", rec.Body.String(), expectedResponse)
	}
}
