package generated

import (
	"context"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
)

func Test___Directive_args_674522442f(t *testing.T) {
	t.Run("Success case", func(t *testing.T) {
		ec := &executionContext{}
		ctx := context.Background()
		field := graphql.CollectedField{}
		obj := &introspection.Directive{
			Args: []introspection.InputValue{
				{Name: "arg1", Description: "desc1"},
			},
		}

		result := ec.___Directive_args(ctx, field, obj)
		if result == graphql.Null {
			t.Errorf("___Directive_args() = %v, want non-null", result)
			t.Logf("Method arguments: ctx = %v, field = %v, obj = %v", ctx, field, obj)
		} else {
			t.Logf("Test success")
		}
	})

	t.Run("Failure case: null result", func(t *testing.T) {
		ec := &executionContext{}
		ctx := context.Background()
		field := graphql.CollectedField{}
		obj := &introspection.Directive{
			Args: nil,
		}

		result := ec.___Directive_args(ctx, field, obj)
		if result != graphql.Null {
			t.Errorf("___Directive_args() = %v, want null", result)
			t.Logf("Method arguments: ctx = %v, field = %v, obj = %v", ctx, field, obj)
		} else {
			t.Logf("Test success")
		}
	})
}
