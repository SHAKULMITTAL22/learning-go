package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestDeleteTask_9ddf916cf4(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/task/1234", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/task/{id}", DeleteTask).Methods("DELETE")

	req = mux.SetURLVars(req, map[string]string{
		"id": "1234",
	})

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("DeleteTask handler returned wrong status code: got %v want %v, id: 1234", status, http.StatusOK)
	}

	expected := `1234`
	if rr.Body.String() != expected {
		t.Errorf("DeleteTask handler returned unexpected body: got %v want %v, id: 1234", rr.Body.String(), expected)
	}

	t.Log("DeleteTask passed successfully.")

	req, err = http.NewRequest("DELETE", "/task/9999", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	req = mux.SetURLVars(req, map[string]string{
		"id": "9999",
	})

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("DeleteTask handler returned wrong status code for non-existent id: got %v want %v, id: 9999", status, http.StatusNotFound)
	}

	t.Log("DeleteTask passed successfully for non-existent task.")
}
