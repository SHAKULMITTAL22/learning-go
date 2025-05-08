// Regenerated Golang test case

package graph

import (
	"context"
	"testing"

	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/database"
	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/graph/model"
)

// Tests the Item function
func TestItem(t *testing.T) {
	db := database.DBConn
	r := queryResolver{db}

	// Test Case 1
	{
		// Mock id
		id := 1
		result, err := r.Item(context.Background(), id)
		if err != nil {
			t.Fatalf("Error encountered: %v", err)
		} else if result.ID != id {
			t.Fatalf("Expected id %v, but got %v ", id, result.ID)
		} 
	}

	// Test Case 2
	{
		// Mock id
		id := 2
		result, err := r.Item(context.Background(), id)
		if err != nil {
			t.Fatalf("Error encountered: %v", err)
		} else if result.ID != id {
			t.Fatalf("Expected id %v, but got %v ", id, result.ID)
		} 
	}
}
