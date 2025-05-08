package generated

import (
	"context"
	"testing"

	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/vektah/gqlparser/v2/ast"
)

type executionContext struct{}

func (ec *executionContext) marshalN__InputValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx context.Context, sel ast.SelectionSet, iv introspection.InputValue) *introspection.InputValue {
	return &iv
}

func TestMarshalN__InputValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(t *testing.T) {
	ec := &executionContext{}

	t.Run("Test with valid input", func(t *testing.T) {
		inputValue := introspection.InputValue{Name: "TestInput", Description: "This is an input value for test", Type: introspection.TypeRef{Name: "String", Kind: "SCALAR"}}
		selectionSet := ast.SelectionSet{}

		result := ec.marshalN__InputValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(context.Background(), selectionSet, inputValue)
		if result == nil {
			t.Error("Error while marshalling: Expected not nil, got nil")
		}
	})

	t.Run("Test with empty InputValue", func(t *testing.T) {
		inputValue := introspection.InputValue{}
		selectionSet := ast.SelectionSet{}

		result := ec.marshalN__InputValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(context.Background(), selectionSet, inputValue)
		if result == nil {
			t.Error("Error while marshalling: Expected not nil, got nil")
		}
	})
}
