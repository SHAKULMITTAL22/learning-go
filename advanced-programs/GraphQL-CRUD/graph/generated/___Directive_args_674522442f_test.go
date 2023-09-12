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

func Test___Directive_args_674522442f(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	field := graphql.CollectedField{
		Name: "__Directive",
		Selections: []*ast.Field{
			{
				Name: "args",
			},
		},
	}

	obj := &introspection.Directive{
		Args: []introspection.InputValue{
			{
				Name: "name",
				Type: gqlparser.String,
			},
			{
				Name: "description",
				Type: gqlparser.String,
			},
		},
	}

	ret := ec.___Directive_args(context.Background(), field, obj)

	if ret == nil {
		t.Error("expected non-nil return value")
	}

	if len(ret.([]introspection.InputValue)) != 2 {
		t.Errorf("expected 2 arguments, got %d", len(ret.([]introspection.InputValue)))
	}

	if ret.([]introspection.InputValue)[0].Name != "name" {
		t.Errorf("expected first argument to be named \"name\", got %s", ret.([]introspection.InputValue)[0].Name)
	}

	if ret.([]introspection.InputValue)[0].Type != gqlparser.String {
		t.Errorf("expected first argument to be of type String, got %s", ret.([]introspection.InputValue)[0].Type)
	}

	if ret.([]introspection.InputValue)[1].Name != "description" {
		t.Errorf("expected second argument to be named \"description\", got %s", ret.([]introspection.InputValue)[1].Name)
	}

	if ret.([]introspection.InputValue)[1].Type != gqlparser.String {
		t.Errorf("expected second argument to be of type String, got %s", ret.([]introspection.InputValue)[1].Type)
	}
}

func Test___Directive_args_with_error(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return nil, errors.New("some error")
		},
	}

	field := graphql.CollectedField{
		Name: "__Directive",
		Selections: []*ast.Field{
			{
				Name: "args",
			},
		},
	}

	obj := &introspection.Directive{
		Args: []introspection.InputValue{
			{
				Name: "name",
				Type: gqlparser.String,
			},
			{
				Name: "description",
				Type: gqlparser.String,
			},
		},
	}

	ret := ec.___Directive_args(context.Background(), field, obj)

	if ret != nil {
		t.Error("expected nil return value")
	}
}
