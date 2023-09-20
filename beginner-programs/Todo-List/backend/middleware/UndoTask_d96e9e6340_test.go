package middleware

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/stretchr/testify/assert"
)

var collection *mongo.Collection

func undoTask(task string) (int64, error) {
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": false}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

func TestUndoTask_d96e9e6340(t *testing.T) {
	// Test case 1: Valid task id
	task1 := "60b5e5c7f790d3a2456874e3"
	t.Log("Executing test case 1 with task id: ", task1)
	modifiedCount, err := undoTask(task1)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), modifiedCount)
	t.Log("Test case 1 passed")

	// Test case 2: Invalid task id
	task2 := "invalid_task_id"
	t.Log("Executing test case 2 with task id: ", task2)
	modifiedCount, err = undoTask(task2)
	assert.NotNil(t, err)
	assert.Equal(t, int64(0), modifiedCount)
	t.Log("Test case 2 passed")
}
