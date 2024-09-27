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

func TestDeleteAllTask_ec6cbb32dc(t *testing.T) {
	// Create a new task
	task := models.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Test Task",
		Description: "This is a test task",
		Completed:   false,
	}

	// Insert the task into the database
	collection.InsertOne(context.Background(), task)

	// Delete all tasks
	deletedCount := deleteAllTask()

	// Assert that the task was deleted
	if deletedCount != 1 {
		t.Errorf("Expected 1 task to be deleted, but got %d", deletedCount)
	}

	// Delete all tasks again
	deletedCount = deleteAllTask()

	// Assert that no tasks were deleted
	if deletedCount != 0 {
		t.Errorf("Expected 0 tasks to be deleted, but got %d", deletedCount)
	}
}
