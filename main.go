package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Server)
	http.HandleFunc("/blogs", Server)
	fmt.Println("Server running on http://localhost:8080 ...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
