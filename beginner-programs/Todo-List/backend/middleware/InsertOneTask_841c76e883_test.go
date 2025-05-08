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

func TestInsertOneTask_841c76e883(t *testing.T) {
	// Create a new task
	task := models.ToDoList{
		Title:       "Test Task",
		Description: "This is a test task",
		Completed:   false,
	}

	// Insert the task into the database
	insertOneTask(task)

	// Get the inserted task from the database
	var insertedTask models.ToDoList
	err := collection.FindOne(context.Background(), bson.M{"_id": task.ID}).Decode(&insertedTask)
	if err != nil {
		t.Error(err)
	}

	// Compare the inserted task with the original task
	if insertedTask.Title != task.Title || insertedTask.Description != task.Description || insertedTask.Completed != task.Completed {
		t.Error("Inserted task does not match original task")
	}
}

func TestInsertOneTask_Error(t *testing.T) {
	// Create a new task with an invalid ID
	task := models.ToDoList{
		ID:          primitive.ObjectID{},
		Title:       "Test Task",
		Description: "This is a test task",
		Completed:   false,
	}

	// Insert the task into the database
	_, err := collection.InsertOne(context.Background(), task)
	if err == nil {
		t.Error("Expected error when inserting task with invalid ID")
	}
}
