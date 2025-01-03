package project

import (
	"encoding/json"
	"net/http"

	"github.com/negagl/myWebsite/helpers"
)

func GetProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(Projects); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func GetProjectByID(w http.ResponseWriter, r *http.Request) {
	// Extract ID
	id, err := helpers.ExtractIDFromPath(w, r, "projects")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Search Project
	project, _, err := FindProjectByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(project); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	// Parse the body
	var newProject Project
	if err := ParseJSONToProject(&newProject, r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Validations
	if err := ValidateProject(newProject); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create Project
	Projects = append(Projects, newProject)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newProject); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	// Extract the ID
	id, err := helpers.ExtractIDFromPath(w, r, "projects")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Parse the body to an object
	var projectToUpdate Project
	if err := ParseJSONToProject(&projectToUpdate, r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Validations
	if err := ValidateProject(projectToUpdate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Search project, update and retrun
	project, _, err := FindProjectByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	project.Title = projectToUpdate.Title
	project.Description = projectToUpdate.Description
	project.URL = projectToUpdate.URL
	project.Status = projectToUpdate.Status

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(project); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	// Extract ID
	id, err := helpers.ExtractIDFromPath(w, r, "projects")
	if err != nil {
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	// Search project
	_, i, err := FindProjectByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Delete project
	Projects = append(Projects[:i], Projects[i+1:]...)

	w.Header().Set("Content-Type", "text/plain")
	if _, err := w.Write([]byte("Project deleted sucessfully\n")); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
