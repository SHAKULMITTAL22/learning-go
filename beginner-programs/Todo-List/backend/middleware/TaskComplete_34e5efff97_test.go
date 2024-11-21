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

func TestTaskComplete_34e5efff97(t *testing.T) {
	// Create a new request
	r, err := http.NewRequest("GET", "/tasks/34e5efff97", nil)
	if err != nil {
		t.Error(err)
	}

	// Create a new response recorder
	w := http.NewRecorder()

	// Create a new context
	ctx := context.Background()

	// Call the TaskComplete function
	TaskComplete(w, r)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Check the response body
	var id string
	if err := json.Unmarshal(w.Body.Bytes(), &id); err != nil {
		t.Error(err)
	}

	if id != "34e5efff97" {
		t.Errorf("Expected task ID %s, got %s", "34e5efff97", id)
	}
}

func TestTaskComplete_InvalidID(t *testing.T) {
	// Create a new request
	r, err := http.NewRequest("GET", "/tasks/invalidID", nil)
	if err != nil {
		t.Error(err)
	}

	// Create a new response recorder
	w := http.NewRecorder()

	// Create a new context
	ctx := context.Background()

	// Call the TaskComplete function
	TaskComplete(w, r)

	// Check the response status code
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}

	// Check the response body
	var errBody models.ErrorBody
	if err := json.Unmarshal(w.Body.Bytes(), &errBody); err != nil {
		t.Error(err)
	}

	if errBody.Error != "Invalid task ID" {
		t.Errorf("Expected error message %s, got %s", "Invalid task ID", errBody.Error)
	}
}
