// Test generated by RoostGPT for test roost-test using AI Type Vertex AI and AI Model code-bison-32k

package main

import (
	"context"
	"fmt"
	"log"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestDeleteAllUsers_cebd155f66(t *testing.T) {
	// Create a new MongoClient and connect to the test database
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Error("Error creating MongoClient:", err)
		return
	}

	// Connect to the test database
	err = client.Connect(context.TODO())
	if err != nil {
		t.Error("Error connecting to the database:", err)
		return
	}

	// Get the users collection
	collection := client.Database("test").Collection("users")

	// Delete all users from the collection
	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		t.Error("Error deleting users from the collection:", err)
		return
	}

	// Check the number of deleted users
	if deleteResult.DeletedCount != 0 {
		t.Errorf("Expected 0 deleted users, got %d", deleteResult.DeletedCount)
		return
	}

	// Print a success message
	fmt.Println("Successfully deleted all users from the collection")
}
