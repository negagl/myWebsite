package main

import (
	"fmt"
	"net/http"
)

func Server(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		w.WriteHeader(http.StatusOK)
		serveContent(w, HomePageContent())
	case "/blogs":
		w.WriteHeader(http.StatusOK)
		serveContent(w, BlogsContent())
	}
	return
}

func serveContent(w http.ResponseWriter, content string) {
	fmt.Fprint(w, content)
}
