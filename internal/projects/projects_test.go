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

func TestGetProjectByID(t *testing.T) {
	tests := []struct {
		name             string
		id               string
		expectedStatus   int
		expectedResponse string
	}{
		{
			"Valid ID",
			"1",
			http.StatusOK,
			`{"id":1,"title":"Project 1","description":"This is the first project","url":"https://github.com/project1","status":"started"}`,
		},
		{
			"Invalid ID",
			"99",
			http.StatusNotFound,
			"Project not found",
		},
		{
			"Invalid ID",
			"wrong",
			http.StatusBadRequest,
			"Invalid project ID",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/projects/"+test.id, nil)
			if err != nil {
				t.Fatalf("Could not create the request: %v", err)
			}

			rec := httptest.NewRecorder()

			GetProjectByID(rec, req)

			if rec.Code != test.expectedStatus {
				t.Errorf("Unexpected status code: got %d, want %d", rec.Code, test.expectedStatus)
			}

			if rec.Body.String() != test.expectedResponse+"\n" {
				t.Errorf("Unexpected respose body: got %q, want %q", rec.Body.String(), test.expectedResponse)
			}
		})
	}
}
