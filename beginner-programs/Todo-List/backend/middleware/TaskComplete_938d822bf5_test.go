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

func TestTaskComplete_938d822bf5(t *testing.T) {
	// Create a new task
	task := models.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Test Task",
		Description: "This is a test task",
		Status:      false,
	}

	// Insert the task into the database
	collection.InsertOne(context.Background(), task)

	// Call the taskComplete function
	taskComplete(task.ID.Hex())

	// Get the updated task from the database
	var updatedTask models.Task
	err := collection.FindOne(context.Background(), bson.M{"_id": task.ID}).Decode(&updatedTask)
	if err != nil {
		t.Error(err)
	}

	// Assert that the task status is true
	if updatedTask.Status != true {
		t.Error("Task status is not true")
	}
}

func TestTaskComplete_InvalidID(t *testing.T) {
	// Call the taskComplete function with an invalid ID
	taskComplete("invalidID")

	// Assert that an error is returned
	if err := collection.FindOne(context.Background(), bson.M{"_id": "invalidID"}).Decode(&models.Task{}); err == nil {
		t.Error("No error returned for invalid ID")
	}
}
