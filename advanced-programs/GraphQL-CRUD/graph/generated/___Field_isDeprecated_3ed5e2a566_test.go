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

func Test___Field_isDeprecated_3ed5e2a566(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	field := graphql.CollectedField{
		Name: "__Field",
		Type: introspection.Field{},
		Args: nil,
	}

	obj := &introspection.Field{
		Name: "isDeprecated",
		Type: "Boolean",
	}

	ctx := context.Background()
	res := ec.___Field_isDeprecated(ctx, field, obj)

	if res == nil {
		t.Error("expected non-nil result")
	}

	if res.(bool) != true {
		t.Errorf("expected true, got %v", res)
	}
}

func Test___Field_isDeprecated_with_error(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return nil, errors.New("some error")
		},
	}

	field := graphql.CollectedField{
		Name: "__Field",
		Type: introspection.Field{},
		Args: nil,
	}

	obj := &introspection.Field{
		Name: "isDeprecated",
		Type: "Boolean",
	}

	ctx := context.Background()
	res := ec.___Field_isDeprecated(ctx, field, obj)

	if res != nil {
		t.Error("expected nil result")
	}

	if !graphql.HasFieldError(ctx, &graphql.FieldContext{
		Object:   "__Field",
		Field:    field,
		Args:     nil,
		IsMethod: true,
	}) {
		t.Error("expected field error")
	}
}
