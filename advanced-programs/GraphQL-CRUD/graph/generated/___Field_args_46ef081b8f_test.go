package generated

import (
	"bytes"
	"context"
	"errors"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

type executionContextTest struct {
	ResolverMiddleware graphql.FieldMiddleware
	Error              func(ctx context.Context, err error)
	Errorf             func(ctx context.Context, format string, args ...interface{})
	Recover            func(ctx context.Context, err interface{}) error
	marshalN__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValueᚄ func(ctx context.Context, sel ast.SelectionSet, v []introspection.InputValue) graphql.Marshaler
}

func Test___Field_args_46ef081b8f(t *testing.T) {
	ec := &executionContextTest{
		ResolverMiddleware: func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
			return next(ctx)
		},
		Error: func(ctx context.Context, err error) {},
		Errorf: func(ctx context.Context, format string, args ...interface{}) {},
		Recover: func(ctx context.Context, err interface{}) error {
			return err.(error)
		},
		marshalN__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValueᚄ: func(ctx context.Context, sel ast.SelectionSet, v []introspection.InputValue) graphql.Marshaler {
			return graphql.Null
		},
	}

	t.Run("Success case", func(t *testing.T) {
		obj := &introspection.Field{
			Args: []introspection.InputValue{},
		}

		field := graphql.CollectedField{}

		result := ec.___Field_args(context.Background(), field, obj)

		if result == graphql.Null {
			t.Error("The method returns null in case of success")
		} else {
			t.Log("Success case passed")
		}
	})

	t.Run("Failure case", func(t *testing.T) {
		obj := &introspection.Field{
			Args: nil,
		}

		field := graphql.CollectedField{}

		result := ec.___Field_args(context.Background(), field, obj)

		if result != graphql.Null {
			t.Errorf("The method should return null in case of failure, got %v", result)
		} else {
			t.Log("Failure case passed")
		}
	})
}
