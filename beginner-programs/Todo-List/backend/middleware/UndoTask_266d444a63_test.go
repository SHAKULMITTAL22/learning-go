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

func TestUndoTask_266d444a63(t *testing.T) {
	// Setup
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Error(err)
	}
	db := client.Database("todo-list")
	collection := db.Collection("tasks")

	// Create a task
	task := models.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Test Task",
		Description: "This is a test task.",
		Completed:   false,
	}
	_, err = collection.InsertOne(ctx, task)
	if err != nil {
		t.Error(err)
	}

	// Undo the task
	w := http.ResponseWriter{}
	r := &http.Request{}
	params := mux.Vars(r)
	params["id"] = task.ID.Hex()
	UndoTask(w, r)

	// Get the task from the database
	var updatedTask models.Task
	err = collection.FindOne(ctx, bson.M{"_id": task.ID}).Decode(&updatedTask)
	if err != nil {
		t.Error(err)
	}

	// Assert that the task is undone
	if updatedTask.Completed {
		t.Error("Task is still completed.")
	}

	// Teardown
	err = collection.DeleteOne(ctx, bson.M{"_id": task.ID})
	if err != nil {
		t.Error(err)
	}
}

func TestUndoTask_Error(t *testing.T) {
	// Setup
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Error(err)
	}
	db := client.Database("todo-list")
	collection := db.Collection("tasks")

	// Create a task
	task := models.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Test Task",
		Description: "This is a test task.",
		Completed:   false,
	}
	_, err = collection.InsertOne(ctx, task)
	if err != nil {
		t.Error(err)
	}

	// Undo the task
	w := http.ResponseWriter{}
	r := &http.Request{}
	params := mux.Vars(r)
	params["id"] = "invalid"
	UndoTask(w, r)

	// Assert that an error was returned
	if w.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.StatusCode)
	}

	// Teardown
	err = collection.DeleteOne(ctx, bson.M{"_id": task.ID})
	if err != nil {
		t.Error(err)
	}
}
