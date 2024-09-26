package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tannergabriel/learning-go/beginner-programs/todo-list/backend/models"
)

// TestCreateTask_eb4d396a39 tests the CreateTask function
func TestCreateTask_eb4d396a39(t *testing.T) {
	// Test case 1: Successful task creation
	t.Run("Success", func(t *testing.T) {
		task := models.ToDoList{
			Task: "Test task",
		}

		requestBody, err := json.Marshal(task)
		if err != nil {
			t.Error(err)
		}

		req, err := http.NewRequest("POST", "/task", bytes.NewBuffer(requestBody))
		if err != nil {
			t.Error(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(CreateTask)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		var createdTask models.ToDoList
		err = json.Unmarshal(rr.Body.Bytes(), &createdTask)
		if err != nil {
			t.Error(err)
		}

		if createdTask.Task != task.Task {
			t.Errorf("handler returned unexpected body: got %v want %v", createdTask.Task, task.Task)
		}

		t.Log("Test case 1: Success")
	})

	// Test case 2: Invalid HTTP method
	t.Run("InvalidMethod", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/task", nil)
		if err != nil {
			t.Error(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(CreateTask)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
		}

		t.Log("Test case 2: InvalidMethod")
	})
}
