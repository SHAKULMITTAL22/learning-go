package middleware

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "testing"

    "github.com/tannergabriel/learning-go/beginner-programs/todo-list/backend/models"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func TestDeleteTask_9ddf916cf4(t *testing.T) {
    // Setup
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    ctx := context.Background()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
    if err != nil {
        log.Fatal("Error connecting to MongoDB")
    }

    db := client.Database(os.Getenv("MONGODB_DATABASE"))
    collection := db.Collection(os.Getenv("MONGODB_COLLECTION"))

    // Create a task
    task := models.Task{
        ID:          primitive.NewObjectID(),
        Title:       "Test Task",
        Description: "This is a test task",
        Completed:   false,
    }

    _, err = collection.InsertOne(ctx, task)
    if err != nil {
        log.Fatal("Error creating task")
    }

    // Delete the task
    req, err := http.NewRequest("DELETE", "/tasks/9ddf916cf4", nil)
    if err != nil {
        t.Fatal(err)
    }

    w := httptest.NewRecorder()
    DeleteTask(w, req)

    // Check the response
    if w.Code != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
    }

    var id string
    err = json.Unmarshal(w.Body.Bytes(), &id)
    if err != nil {
        t.Fatal(err)
    }

    if id != "9ddf916cf4" {
        t.Errorf("Expected task ID %s, got %s", "9ddf916cf4", id)
    }

    // Cleanup
    _, err = collection.DeleteOne(ctx, bson.M{"_id": task.ID})
    if err != nil {
        log.Fatal("Error deleting task")
    }
}
