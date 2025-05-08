// Generated Test case by AI Model

package generated

import (
	"context"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/graph/model"
)

// TestUnmarshalInputNewItem_a630e54bf7 is a test case for the method unmarshalInputNewItem
func TestUnmarshalInputNewItem_a630e54bf7(t *testing.T) {
	gotch := New()
	ctx := context.Background()

	t.Log("Testing unmarshalInputNewItem with valid inputs")

	// Test with valid inputs
	input := map[string]interface{}{"title": "A demo title", "owner": "John Doe", "rating": 5}
	expectedOutput := model.NewItem{Title: "A demo title", Owner: "John Doe", Rating: 5}

	result, err := gotch.unmarshalInputNewItem(ctx, input)
	if err != nil {
		t.Errorf("Failed while testing unmarshalInputNewItem with valid inputs, returned error: %v", err)
	}

	if result.Title != expectedOutput.Title || result.Owner != expectedOutput.Owner || result.Rating != expectedOutput.Rating {
		t.Errorf("Failed, expected %v, but got %v", expectedOutput, result)
	}

	t.Log("Test passed. unmarshalInputNewItem works correctly with valid inputs")

	t.Log("Testing unmarshalInputNewItem with missing inputs")

	// Test with missing inputs
	input2 := map[string]interface{}{"title": "Another demo title", "rating": 10}

	_, err = gotch.unmarshalInputNewItem(ctx, input2)
	if err == nil || err.Error() != "object key owner is missing" {
		t.Errorf("Failed while testing unmarshalInputNewItem with missing inputs. Method did not return expected error.")
	}

	t.Log("Test passed. unmarshalInputNewItem correctly handles missing inputs")
}
