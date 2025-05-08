package generated

import (
	"context"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
)

type executionContext struct{}

func (ec *executionContext) ___Directive_name(ctx context.Context, field graphql.CollectedField, directive *introspection.Directive) *string {
	if directive == nil {
		return nil
	}
	return &directive.Name
}

func Test___Directive_name(t *testing.T) {
	ec := &executionContext{}
	field := graphql.CollectedField{}

	t.Run("Test with valid directive", func(t *testing.T) {
		directive := &introspection.Directive{Name: "validDirective"}
		result := ec.___Directive_name(context.Background(), field, directive)

		if result == nil {
			t.Error("Expected non-nil result, got nil")
			t.Logf("Directive: %v", directive)
		} else {
			t.Log("Success: Test with valid directive")
		}
	})

	t.Run("Test with nil directive", func(t *testing.T) {
		result := ec.___Directive_name(context.Background(), field, nil)

		if result != nil {
			t.Error("Expected nil result, got non-nil")
			t.Log("Success: Test with nil directive")
		}
	})

	t.Run("Test with empty directive name", func(t *testing.T) {
		directive := &introspection.Directive{Name: ""}
		result := ec.___Directive_name(context.Background(), field, directive)

		if result == nil {
			t.Error("Expected non-nil result, got nil")
			t.Logf("Directive: %v", directive)
		} else {
			t.Log("Success: Test with empty directive name")
		}
	})
}
