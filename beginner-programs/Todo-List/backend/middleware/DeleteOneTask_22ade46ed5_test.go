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

func TestDeleteOneTask_22ade46ed5(t *testing.T) {
	// Create a new task
	task := models.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Test Task",
		Description: "This is a test task",
		Completed:   false,
	}

	// Insert the task into the database
	collection.InsertOne(context.Background(), task)

	// Delete the task
	deleteOneTask(task.ID.Hex())

	// Check that the task was deleted
	filter := bson.M{"_id": task.ID}
	var deletedTask models.Task
	err := collection.FindOne(context.Background(), filter).Decode(&deletedTask)
	if err != nil {
		t.Error("Error finding task:", err)
	}

	if deletedTask.ID != "" {
		t.Error("Task was not deleted")
	}
}

func TestDeleteOneTask_InvalidID(t *testing.T) {
	// Delete a task with an invalid ID
	deleteOneTask("invalidID")

	// Check that the task was not deleted
	filter := bson.M{"_id": "invalidID"}
	var deletedTask models.Task
	err := collection.FindOne(context.Background(), filter).Decode(&deletedTask)
	if err != nil {
		t.Error("Error finding task:", err)
	}

	if deletedTask.ID != "" {
		t.Error("Task was not deleted")
	}
}
