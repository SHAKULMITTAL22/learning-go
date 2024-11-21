package generated

import (
	"bytes"
	"context"
	"errors"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/graph/model"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

func Test___Field_args_46ef081b8f(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	field := graphql.CollectedField{
		Name: "__Field",
		Selections: []*ast.Field{
			{
				Name: "args",
			},
		},
	}

	obj := &introspection.Field{
		Args: []introspection.InputValue{
			{
				Name: "name",
				Type: &introspection.Type{
					Kind: "String",
				},
			},
			{
				Name: "age",
				Type: &introspection.Type{
					Kind: "Int",
				},
			},
		},
	}

	ret := ec.___Field_args(context.Background(), field, obj)

	if ret == nil {
		t.Error("expected non-nil result")
	}

	if len(ret.([]introspection.InputValue)) != 2 {
		t.Errorf("expected 2 args, got %d", len(ret.([]introspection.InputValue)))
	}

	if ret.([]introspection.InputValue)[0].Name != "name" {
		t.Errorf("expected first arg to be named 'name', got '%s'", ret.([]introspection.InputValue)[0].Name)
	}

	if ret.([]introspection.InputValue)[1].Name != "age" {
		t.Errorf("expected second arg to be named 'age', got '%s'", ret.([]introspection.InputValue)[1].Name)
	}
}

func Test___Field_args_with_error(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return nil, errors.New("some error")
		},
	}

	field := graphql.CollectedField{
		Name: "__Field",
		Selections: []*ast.Field{
			{
				Name: "args",
			},
		},
	}

	obj := &introspection.Field{
		Args: []introspection.InputValue{
			{
				Name: "name",
				Type: &introspection.Type{
					Kind: "String",
				},
			},
			{
				Name: "age",
				Type: &introspection.Type{
					Kind: "Int",
				},
			},
		},
	}

	ret := ec.___Field_args(context.Background(), field, obj)

	if ret != nil {
		t.Error("expected nil result")
	}
}
