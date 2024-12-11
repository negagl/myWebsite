package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "ok"}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	jsonResponse, _ := json.Marshal(response)
	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/health", healthCheckHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to my personal website!")
	})

	fmt.Println("Server is running on http://localhost:8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("There was an error starting the server")
		return
	}
}
