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
