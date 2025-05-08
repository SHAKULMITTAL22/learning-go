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

func TestGetAllTask_dfbd966f8b(t *testing.T) {
    // Setup
    ctx := context.Background()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        t.Fatal(err)
    }
    collection := client.Database("todo-list").Collection("tasks")

    // Create a few tasks
    tasks := []models.Task{
        {ID: primitive.NewObjectID(), Name: "Task 1"},
        {ID: primitive.NewObjectID(), Name: "Task 2"},
        {ID: primitive.NewObjectID(), Name: "Task 3"},
    }
    for _, task := range tasks {
        _, err := collection.InsertOne(ctx, task)
        if err != nil {
            t.Fatal(err)
        }
    }

    // Call the method under test
    results := getAllTask()

    // Check the results
    if len(results) != len(tasks) {
        t.Errorf("Expected %d results, got %d", len(tasks), len(results))
    }
    for i, result := range results {
        if result["name"] != tasks[i].Name {
            t.Errorf("Expected task %d to have name %s, got %s", i, tasks[i].Name, result["name"])
        }
    }

    // Cleanup
    for _, task := range tasks {
        _, err := collection.DeleteOne(ctx, bson.M{"_id": task.ID})
        if err != nil {
            t.Fatal(err)
        }
    }
}

func TestGetAllTask_empty(t *testing.T) {
    // Setup
    ctx := context.Background()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        t.Fatal(err)
    }
    collection := client.Database("todo-list").Collection("tasks")

    // Call the method under test
    results := getAllTask()

    // Check the results
    if len(results) != 0 {
        t.Errorf("Expected 0 results, got %d", len(results))
    }

    // Cleanup
}
