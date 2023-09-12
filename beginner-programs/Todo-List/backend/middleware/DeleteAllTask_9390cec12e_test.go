package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tannergabriel/learning-go/beginner-programs/todo-list/backend/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestDeleteAllTask_9390cec12e(t *testing.T) {
	// Test case 1: Success
	req, err := http.NewRequest("DELETE", "/tasks", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteAllTask)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	expected := models.Count{Count: 1}
	var actual models.Count
	err = json.Unmarshal(rr.Body.Bytes(), &actual)
	if err != nil {
		t.Error(err)
	}

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			actual, expected)
	}

	// Test case 2: Failure
	req, err = http.NewRequest("DELETE", "/tasks", nil)
	if err != nil {
		t.Error(err)
	}

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(DeleteAllTask)

	// Simulate a database error
	models.DB.Client().Disconnect(context.TODO())

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

	// Check the response body
	expected = models.Count{Count: 0}
	err = json.Unmarshal(rr.Body.Bytes(), &actual)
	if err != nil {
		t.Error(err)
	}

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			actual, expected)
	}
}
