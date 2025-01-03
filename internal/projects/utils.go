package projects

import (
	"encoding/json"
	"errors"
	"net/http"
)

func FindProjectByID(id int) (*Project, int, error) {
	for i, project := range Projects {
		if project.ID == id {
			return &project, i, nil
		}
	}

	return nil, -1, errors.New("Project not found")
}

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

func ParseJSONToProject(projectToParse *Project, r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(&projectToParse); err != nil {
		return errors.New("Invalid JSON")
	}

	return nil
}
