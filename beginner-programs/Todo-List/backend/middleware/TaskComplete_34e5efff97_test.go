package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// TestTaskComplete_34e5efff97 is a test function for the TaskComplete function
func TestTaskComplete_34e5efff97(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/task/{id}", TaskComplete).Methods("POST")

	// Test case 1: Valid task ID
	req, err := http.NewRequest("POST", "/task/123", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("TaskComplete handler returned wrong status code: got %v want %v", status, http.StatusOK)
		t.Log("Method Arguments: ", req.URL.String())
	} else {
		t.Log("Success: Valid Task ID")
	}

	// Test case 2: Invalid task ID
	req, err = http.NewRequest("POST", "/task/abc", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("TaskComplete handler returned wrong status code for invalid ID: got %v want %v", status, http.StatusBadRequest)
		t.Log("Method Arguments: ", req.URL.String())
	} else {
		t.Log("Success: Invalid Task ID handled correctly")
	}

	// Test case 3: No task ID
	req, err = http.NewRequest("POST", "/task/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("TaskComplete handler returned wrong status code for missing ID: got %v want %v", status, http.StatusBadRequest)
		t.Log("Method Arguments: ", req.URL.String())
	} else {
		t.Log("Success: Missing Task ID handled correctly")
	}
}
