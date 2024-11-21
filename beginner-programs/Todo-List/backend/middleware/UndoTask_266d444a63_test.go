package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// TestUndoTask_266d444a63 is a test function for UndoTask function
func TestUndoTask_266d444a63(t *testing.T) {
	// Create a request to pass to our handler.
	req, err := http.NewRequest("GET", "/task/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/task/{id}", UndoTask).Methods("GET")

	// Create a request to pass to our handler.
	router.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body is what we expect.
	expected := `"123"`
	assert.Equal(t, expected, rr.Body.String())

	// Test case when the id doesn't exist
	req, err = http.NewRequest("GET", "/task/999", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body is what we expect.
	expected = `Task not found`
	assert.Contains(t, rr.Body.String(), expected)
}
