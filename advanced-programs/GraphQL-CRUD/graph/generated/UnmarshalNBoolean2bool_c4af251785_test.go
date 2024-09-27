// Test created by RoostGPT for test azure-32k-go using AI Type Azure Open AI and AI Model roost-gpt4-32k

package generated

import (
	"context"
	"testing"

	"github.com/99designs/gqlgen/graphql"
)

type executionContext struct{}

func (ec *executionContext) unmarshalNBoolean2bool(ctx context.Context, v interface{}) (bool, error) {
	return graphql.UnmarshalBoolean(v)
}

// TestUnmarshalNBoolean2bool verifies the unmarshalNBoolean2bool operation
func TestUnmarshalNBoolean2bool(t *testing.T) {
	// create test scenarios
	scenarios := []struct {
		input    interface{}
		expected bool
		isError  bool
	}{
		{input: true, expected: true, isError: false},
		{input: false, expected: false, isError: false},
		{input: "true", isError: true},
		{input: "false", isError: true},
		{input: 123, isError: true},
	}

	// create instance of executionContext
	ctx := context.Background()
	ec := &executionContext{}

	// process test scenarios
	for _, s := range scenarios {
		outcome, err := ec.unmarshalNBoolean2bool(ctx, s.input)

		if (err != nil) != s.isError {
			t.Fatalf("Unexpected error for input %v. Error: %v", s.input, err)
		}

		if outcome != s.expected {
			t.Errorf("Unexpected outcome for input %v. Want: %v, Got: %v", s.input, s.expected, outcome)
		}
	}
}
