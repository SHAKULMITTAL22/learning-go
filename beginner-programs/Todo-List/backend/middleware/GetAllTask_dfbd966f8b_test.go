package middleware

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

// Mock data for testing
var mockData = []primitive.M{
	{"_id": primitive.NewObjectID(), "task": "Test task 1", "status": "pending"},
	{"_id": primitive.NewObjectID(), "task": "Test task 2", "status": "completed"},
}

// TestGetAllTask_dfbd966f8b is a test function for the getAllTask function.
func TestGetAllTask_dfbd966f8b(t *testing.T) {
	// TODO: Replace with your MongoDB connection string
	clientOptions := options.Client().ApplyURI("<YOUR_MONGODB_CONNECTION_STRING>")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		t.Error(err)
	}
	collection = client.Database("testDB").Collection("testCollection")

	// Insert mock data into the collection
	_, err = collection.InsertMany(context.Background(), mockData)
	if err != nil {
		t.Error(err)
	}

	// Test case 1: Check if getAllTask returns all tasks
	tasks := getAllTask()
	if len(tasks) != len(mockData) {
		t.Errorf("getAllTask() failed, expected %v, got %v", len(mockData), len(tasks))
		t.Log("getAllTask arguments: None")
	} else {
		t.Log("Success: getAllTask returned all tasks")
	}

	// Test case 2: Check if getAllTask returns correct tasks
	for i, task := range tasks {
		if task != mockData[i] {
			t.Errorf("getAllTask() failed, expected %v, got %v", mockData[i], task)
			t.Log("getAllTask arguments: None")
		} else {
			t.Log("Success: getAllTask returned correct tasks")
		}
	}

	// Cleanup: Remove mock data from the collection
	_, err = collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		t.Error(err)
	}
}
