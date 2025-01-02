package projects

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func GetProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(projects); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func GetProjectByID(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Path[len("/projects/"):]
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	for _, project := range projects {
		if project.ID == id {
			w.WriteHeader(http.StatusOK)
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
	var newProject Project
	if err := json.NewDecoder(r.Body).Decode(&newProject); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

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
	idString := r.URL.Path[len("/projects/"):]
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
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

	// Look for the project, update and retrun
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
	idString := r.URL.Path[len("/projects/"):]
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, project := range projects {
		if project.ID == id {
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
