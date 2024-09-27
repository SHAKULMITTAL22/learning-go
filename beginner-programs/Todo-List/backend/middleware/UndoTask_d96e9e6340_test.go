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

func TestUndoTask_d96e9e6340(t *testing.T) {
	// Create a new task
	task := models.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Test Task",
		Description: "This is a test task",
		Status:      true,
	}

	// Insert the task into the database
	collection.InsertOne(context.Background(), task)

	// Call the undoTask function
	undoTask(task.ID.Hex())

	// Get the updated task from the database
	var updatedTask models.Task
	err := collection.FindOne(context.Background(), bson.M{"_id": task.ID}).Decode(&updatedTask)
	if err != nil {
		t.Error(err)
	}

	// Assert that the task status is false
	if updatedTask.Status != false {
		t.Error("Task status is not false")
	}
}

func TestUndoTask_InvalidID(t *testing.T) {
	// Call the undoTask function with an invalid ID
	undoTask("invalidID")

	// Assert that an error was returned
	if err := recover(); err == nil {
		t.Error("No error was returned")
	}
}
