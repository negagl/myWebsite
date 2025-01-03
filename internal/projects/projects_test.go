package projects

import (
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestCreateProject(t *testing.T) {
	tests := []struct {
		name             string
		body             string
		expectedStatus   int
		expectedResponse string
	}{
		{
			name:             "Project created correctly",
			body:             `{"id":2,"title":"New Project","description":"This is a new project","url":"https://github.com/newproject","status":"created"}`,
			expectedStatus:   http.StatusCreated,
			expectedResponse: `{"id":2,"title":"New Project","description":"This is a new project","url":"https://github.com/newproject","status":"created"}`,
		},
		{
			name:             "Empty title",
			body:             `{"id":3,"title":"","description":"This is a new project","url":"https://github.com/newproject","status":"created"}`,
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: "Title cannot be empty",
		},
		{
			name:             "Duplicated ID",
			body:             `{"id":1,"title":"Existing Project","description":"This is a existing project","url":"https://github.com/newproject","status":"created"}`,
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: "ID already exists",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/projects", strings.NewReader(test.body))
			if err != nil {
				t.Fatalf("Could not create the request: %v", err)
			}

			rec := httptest.NewRecorder()

			CreateProject(rec, req)

			if rec.Code != test.expectedStatus {
				t.Errorf("Unexpected status code: got %d, want %d", rec.Code, test.expectedStatus)
			}

			if rec.Body.String() != test.expectedResponse+"\n" {
				t.Errorf("Unexpected response body: got %q, want %q", rec.Body.String(), test.expectedResponse)
			}
		})
	}
}

func TestUpdateProject(t *testing.T) {
	tests := []struct {
		name             string
		id               string
		body             string
		expectedStatus   int
		expectedResponse string
	}{
		{
			name:             "Valid ID",
			id:               "1",
			body:             `{"title":"New title","description":"New description","url":"","status":"on hold"}`,
			expectedStatus:   http.StatusOK,
			expectedResponse: `{"id":1,"title":"New title","description":"New description","url":"","status":"on hold"}`,
		},
		{
			name:             "ID doesnt exists",
			id:               "99",
			body:             `{"title":"New title","description":"New description","url":"","status":"on hold"}`,
			expectedStatus:   http.StatusNotFound,
			expectedResponse: "Invalid ID",
		},
		{
			name:             "Empty Title",
			id:               "1",
			body:             `{"title":"","description":"New description","url":"","status":"on hold"}`,
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: "Title cannot be empty",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPut, "/projects/"+test.id, strings.NewReader(test.body))
			if err != nil {
				t.Fatalf("Could not create teh request: %v", err)
			}

			rec := httptest.NewRecorder()

			UpdateProject(rec, req)

			if rec.Code != test.expectedStatus {
				t.Errorf("Unexpected status code: got %d, want %d", rec.Code, test.expectedStatus)
			}

			if rec.Body.String() != test.expectedResponse+"\n" {
				t.Errorf("Unexpected response body: got %q, want %q", rec.Body.String(), test.expectedResponse)
			}
		})
	}
}

func TestDeleteProject(t *testing.T) {
	tests := []struct {
		name             string
		id               string
		expectedStatus   int
		expectedResponse string
	}{
		{
			name:             "Deleted Succesfully",
			id:               "1",
			expectedStatus:   http.StatusOK,
			expectedResponse: "Project deleted sucessfully",
		},
		{
			name:             "ID not found",
			id:               "99",
			expectedStatus:   http.StatusNotFound,
			expectedResponse: "Project not found",
		},
		{
			name:             "Invalid ID",
			id:               "a",
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: "Invalid project ID",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/projects/"+test.id, nil)
			if err != nil {
				t.Fatalf("Could not create the request: %v", err)
			}

			rec := httptest.NewRecorder()

			DeleteProject(rec, req)

			if rec.Code != test.expectedStatus {
				t.Errorf("Unexpected status code: got %d, want %d", rec.Code, test.expectedStatus)
			}

			if test.expectedResponse != "" && rec.Body.String() != test.expectedResponse+"\n" {
				t.Errorf("Unexpected response body: got %q, want %q", rec.Body.String(), test.expectedResponse)
			}
		})
	}
}
