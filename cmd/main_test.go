package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckHealth(t *testing.T) {
	// Step 1: Create and HTTP request
	req, err := http.NewRequest(http.MethodGet, "/health", nil)
	if err != nil {
		t.Fatalf("Could not create the request: %v", err)
	}

	// Step 2: Record the response
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(healthCheckHandler)

	// Step 3: Call the handler
	handler.ServeHTTP(rec, req)

	// Step 4: Assert the response
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, rec.Code)
	}

	expected := `{"status":"ok"}`
	if rec.Body.String() != expected {
		t.Errorf("Unexpected body: got %v, want %v", rec.Body.String(), expected)
	}
}
