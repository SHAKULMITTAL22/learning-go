package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestDeleteAllTask_9390cec12e is a test function for DeleteAllTask function
func TestDeleteAllTask_9390cec12e(t *testing.T) {
	req, err := http.NewRequest("GET", "/deleteAllTask", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteAllTask)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"count":0}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	t.Log("DeleteAllTask passed")
}

// TestDeleteAllTask_9390cec12e_Fail is a test function for DeleteAllTask function when it fails
func TestDeleteAllTask_9390cec12e_Fail(t *testing.T) {
	req, err := http.NewRequest("GET", "/deleteAllTask", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteAllTask)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	t.Log("DeleteAllTask failed as expected")
}
