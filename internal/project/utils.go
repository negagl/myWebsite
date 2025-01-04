package project

import (
	"encoding/json"
	"errors"
	"net/http"
)

// FindProjectByID finds a project by its ID
func FindProjectByID(id int) (*Project, int, error) {
	for i, project := range Projects {
		if project.ID == id {
			return &Projects[i], i, nil
		}
	}

	return nil, -1, errors.New("Project not found")
}

// ValidateProject checks if a project passed is valid to the structure defined
func ValidateProject(project Project) error {
	if project.Title == "" {
		return errors.New("Title cannot be empty")
	}

	for _, pjt := range Projects {
		if pjt.ID == project.ID {
			return errors.New("ID already exists")
		}
	}

	return nil
}

// ParseJSONToProject takes a json body from the request and parse it to a project
func ParseJSONToProject(projectToParse *Project, r *http.Request) error {
	if r.Body == nil {
		return errors.New("Request body cannot be empty")
	}

	if err := json.NewDecoder(r.Body).Decode(&projectToParse); err != nil {
		return errors.New("Invalid JSON")
	}

	return nil
}
