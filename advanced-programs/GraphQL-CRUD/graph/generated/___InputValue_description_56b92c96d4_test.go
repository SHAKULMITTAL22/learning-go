package generated

import (
	"context"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
)

// Testing function ___InputValue_description.
func Test___InputValue_description_56b92c96d4(t *testing.T) {
	execContext := &ExecutionContext{}
	collectedField := graphql.CollectedField{}
	inputValue := &introspection.InputValue{Description: "Test Description"}

	// Test case: Success scenario
	t.Run("Success Scenario", func(t *testing.T) {
		inputValue.Description = "Test description" // TODO: Change to the required description
		result := execContext.___InputValue_description(context.Background(), &collectedField, inputValue)
		if result == nil {
			t.Errorf("Expected non-nil result")
		} else {
			t.Logf("Test___InputValue_description_56b92c96d4 Test case: Success scenario PASSED")
		}
	})

	// Test case: Failure scenario - When inputValue is nil
	t.Run("Failure Scenario", func(t *testing.T) {
		result := execContext.___InputValue_description(context.Background(), &collectedField, nil)
		if result != nil {
			t.Errorf("Expected nil result ")
		} else {
			t.Logf("Test___InputValue_description_56b92c96d4 Test case: Failure scenario PASSED")
		}
	})

	// Test case: Edge scenario - When Description is empty
	t.Run("Edge Scenario", func(t *testing.T) {
		inputValue.Description = ""
		result := execContext.___InputValue_description(context.Background(), &collectedField, inputValue)
		if result == nil {
			t.Errorf("Expected non-nil result ")
		} else {
			t.Logf("Test___InputValue_description_56b92c96d4 Test case: Edge scenario PASSED")
		}
	})
}
