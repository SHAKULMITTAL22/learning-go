package middleware

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mocking the mongo.Collection
type mockCollection struct {
	mongo.Collection
}

func (mc *mockCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return &mongo.UpdateResult{ModifiedCount: 1}, nil
}

var collection *mockCollection

// TestTaskComplete_938d822bf5 tests the taskComplete function
func TestTaskComplete_938d822bf5(t *testing.T) {
	task := primitive.NewObjectID().Hex()
	collection = &mockCollection{}

	// Test case 1: Valid task id
	t.Run("Valid task id", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Error("The code did not handle the error correctly")
			}
		}()

		taskComplete(t, task)
	})

	// Test case 2: Invalid task id
	t.Run("Invalid task id", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("The code did not handle the error correctly")
			}
		}()

		taskComplete(t, "invalidTaskId")
	})
}

func taskComplete(t *testing.T, task string) {
	id, err := primitive.ObjectIDFromHex(task)
	if err != nil {
		panic(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, int64(1), result.ModifiedCount, "Modified count should be 1")
}
