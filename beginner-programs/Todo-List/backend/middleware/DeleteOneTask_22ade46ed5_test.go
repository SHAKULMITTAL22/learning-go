package middleware

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tannergabriel/learning-go/beginner-programs/todo-list/backend/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection // TODO: Initialize this with your collection

// Test case for deleteOneTask method
func TestDeleteOneTask_22ade46ed5(t *testing.T) {
	// Test case 1: Valid task ID
	taskID := "60d5ecf9e1f6133dd8b062e7" // TODO: Replace with a valid task ID
	t.Run("valid id", func(t *testing.T) {
		// Call the method with a valid task ID
		deleteOneTask(taskID)

		// Check if the task was deleted
		id, _ := primitive.ObjectIDFromHex(taskID)
		filter := bson.M{"_id": id}
		err := collection.FindOne(context.Background(), filter).Err()
		if err != mongo.ErrNoDocuments {
			t.Error("Task was not deleted")
			t.Log("Task ID:", taskID)
		} else {
			t.Log("Success: Task was deleted")
		}
	})

	// Test case 2: Invalid task ID
	taskID = "invalidID"
	t.Run("invalid id", func(t *testing.T) {
		// Call the method with an invalid task ID
		defer func() {
			if r := recover(); r == nil {
				t.Error("The code did not panic")
				t.Log("Task ID:", taskID)
			} else {
				t.Log("Success: The code panicked as expected")
			}
		}()
		deleteOneTask(taskID)
	})
}

func deleteOneTask(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	d, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
}
