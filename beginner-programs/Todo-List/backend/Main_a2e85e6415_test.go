package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tannergabriel/learning-go/beginner-programs/todo-list/backend/router"
)

// TestMain_a2e85e6415 is a test function for the main function
func TestMain_a2e85e6415(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/todos", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()
	handler := router.Router()

	// Call ServeHTTP method directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"todos":[]}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	t.Log("TestMain_a2e85e6415 passed successfully")
}
