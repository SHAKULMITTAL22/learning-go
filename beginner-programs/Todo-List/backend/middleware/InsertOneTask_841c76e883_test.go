package middleware

import (
	"context"
	"testing"
	"log"
	"fmt"

	"github.com/tannergabriel/learning-go/beginner-programs/todo-list/backend/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func insertOneTask(task models.ToDoList) {
	insertResult, err := collection.InsertOne(context.Background(), task)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}

func TestInsertOneTask_841c76e883(t *testing.T) {
	task1 := models.ToDoList{ID: primitive.NewObjectID(), Task: "Task 1", Status: false}
	task2 := models.ToDoList{ID: primitive.NewObjectID(), Task: "Task 2", Status: true}

	t.Run("insert task1", func(t *testing.T) {
		assert.NotPanics(t, func() { insertOneTask(task1) }, "The code did not panic")
		t.Log("Success: Task 1 inserted successfully")
	})

	t.Run("insert task2", func(t *testing.T) {
		assert.NotPanics(t, func() { insertOneTask(task2) }, "The code did not panic")
		t.Log("Success: Task 2 inserted successfully")
	})

	t.Run("insert same task again", func(t *testing.T) {
		assert.Panics(t, func() { insertOneTask(task1) }, "The code did panic")
		t.Log("Failure: Same task cannot be inserted again")
	})
}
