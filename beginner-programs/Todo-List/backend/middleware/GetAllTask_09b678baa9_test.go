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

func TestGetAllTask_09b678baa9(t *testing.T) {
	// Setup
	ctx := context.Background()
	godotenv.Load()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		t.Error(err)
	}
	defer client.Disconnect(ctx)
	db := client.Database("todo_list")
	collection := db.Collection("tasks")

	// Create a task
	task := models.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Test Task",
		Description: "This is a test task",
		Completed:   false,
	}
	_, err = collection.InsertOne(ctx, task)
	if err != nil {
		t.Error(err)
	}

	// Create a request
	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Error(err)
	}

	// Create a response recorder
	w := httptest.NewRecorder()

	// Call the handler
	GetAllTask(w, req)

	// Check the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Decode the response body
	var tasks []models.Task
	err = json.Unmarshal(w.Body.Bytes(), &tasks)
	if err != nil {
		t.Error(err)
	}

	// Check the response body
	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}

	if tasks[0].ID != task.ID {
		t.Errorf("Expected task ID %s, got %s", task.ID, tasks[0].ID)
	}

	if tasks[0].Title != task.Title {
		t.Errorf("Expected task title %s, got %s", task.Title, tasks[0].Title)
	}

	if tasks[0].Description != task.Description {
		t.Errorf("Expected task description %s, got %s", task.Description, tasks[0].Description)
	}

	if tasks[0].Completed != task.Completed {
		t.Errorf("Expected task completed %t, got %t", task.Completed, tasks[0].Completed)
	}

	// Cleanup
	_, err = collection.DeleteOne(ctx, bson.M{"_id": task.ID})
	if err != nil {
		t.Error(err)
	}
}

func TestGetAllTask_NoTasks(t *testing.T) {
	// Setup
	ctx := context.Background()
	godotenv.Load()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		t.Error(err)
	}
	defer client.Disconnect(ctx)
	db := client.Database("todo_list")
	collection := db.Collection("tasks")

	// Create a request
	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Error(err)
	}

	// Create a response recorder
	w := httptest.NewRecorder()

	// Call the handler
	GetAllTask(w, req)

	// Check the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Decode the response body
	var tasks []models.Task
	err = json.Unmarshal(w.Body.Bytes(), &tasks)
	if err != nil {
		t.Error(err)
	}

	// Check the response body
	if len(tasks) != 0 {
		t.Errorf("Expected 0 tasks, got %d", len(tasks))
	}

	// Cleanup
	_, err = collection.DeleteOne(ctx, bson.M{})
	if err != nil {
		t.Error(err)
	}
}
