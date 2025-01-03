package projects

import (
	"encoding/json"
	"net/http"

	"github.com/negagl/myWebsite/utils"
)

func GetProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(projects); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func GetProjectByID(w http.ResponseWriter, r *http.Request) {
	// Extract ID
	id, err := utils.ExtractIDFromPath(w, r, "projects")
	if err != nil {
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	// Search Project
	for _, project := range projects {
		if project.ID == id {
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(project); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			return
		}
	}

	http.Error(w, "Project not found", http.StatusNotFound)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	// Parse the body
	var newProject Project
	if err := json.NewDecoder(r.Body).Decode(&newProject); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validations
	if newProject.Title == "" {
		http.Error(w, "Title cannot be empty", http.StatusBadRequest)
		return
	}

	for _, project := range projects {
		if project.ID == newProject.ID {
			http.Error(w, "ID already exists", http.StatusBadRequest)
			return
		}
	}

	// Create Project
	projects = append(projects, newProject)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newProject); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	// Extract the ID
	id, err := utils.ExtractIDFromPath(w, r, "projects")
	if err != nil {
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	// Parse the body to an object
	var projectToUpdate Project
	if err := json.NewDecoder(r.Body).Decode(&projectToUpdate); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validations
	if projectToUpdate.Title == "" {
		http.Error(w, "Title cannot be empty", http.StatusBadRequest)
		return
	}

	// Search project, update and retrun
	for i, project := range projects {
		if project.ID == id {
			projects[i].Title = projectToUpdate.Title
			projects[i].Description = projectToUpdate.Description
			projects[i].URL = projectToUpdate.URL
			projects[i].Status = projectToUpdate.Status

			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(projects[i]); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			return
		}
	}

	http.Error(w, "Invalid ID", http.StatusNotFound)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	// Extract ID
	id, err := utils.ExtractIDFromPath(w, r, "projects")
	if err != nil {
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	// Search project
	for i, project := range projects {
		if project.ID == id {

			// Delete project
			projects = append(projects[:i], projects[i+1:]...)

			if _, err := w.Write([]byte("Project deleted sucessfully\n")); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			return
		}
	}

	http.Error(w, "Project not found", http.StatusNotFound)
}
