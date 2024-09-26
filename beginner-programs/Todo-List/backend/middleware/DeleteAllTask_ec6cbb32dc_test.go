package middleware

import (
	"context"
	"errors"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mocking the MongoDB collection
type MockCollection struct {
	mockDeleteCount int64
	mockDeleteErr   error
}

func (mc *MockCollection) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return &mongo.DeleteResult{DeletedCount: mc.mockDeleteCount}, mc.mockDeleteErr
}

// Mocking deleteAllTask function
func deleteAllTask(collection *MockCollection) int64 {
	deleteResult, _ := collection.DeleteMany(context.Background(), bson.D{{}})
	return deleteResult.DeletedCount
}

// Test case for deleteAllTask function
func TestDeleteAllTask_ec6cbb32dc(t *testing.T) {
	tests := []struct {
		name            string
		mockDeleteCount int64
		mockDeleteErr   error
		want            int64
	}{
		{
			name:            "Test case 1: When deletion is successful",
			mockDeleteCount: 5,
			mockDeleteErr:   nil,
			want:            5,
		},
		{
			name:            "Test case 2: When deletion fails",
			mockDeleteCount: 0,
			mockDeleteErr:   errors.New("deletion error"),
			want:            0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collection := &MockCollection{
				mockDeleteCount: tt.mockDeleteCount,
				mockDeleteErr:   tt.mockDeleteErr,
			}
			if got := deleteAllTask(collection); got != tt.want {
				t.Errorf("deleteAllTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
