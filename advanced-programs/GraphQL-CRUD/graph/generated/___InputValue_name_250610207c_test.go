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

func Test___InputValue_name_250610207c(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	// Success case
	obj := &introspection.InputValue{
		Name: "name",
	}
	field := graphql.CollectedField{
		Selections: []*ast.Field{
			{
				Name: "name",
			},
		},
	}
	ctx := context.Background()
	res := ec.___InputValue_name(ctx, field, obj)
	if res == nil {
		t.Error("expected non-nil result")
	}
	if res.(string) != "name" {
		t.Errorf("expected name to be 'name', got '%s'", res.(string))
	}

	// Error case
	obj = &introspection.InputValue{
		Name: "",
	}
	res = ec.___InputValue_name(ctx, field, obj)
	if res != nil {
		t.Error("expected nil result")
	}
	if !graphql.HasFieldError(ctx, &graphql.FieldContext{
		Object:   "__InputValue",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}) {
		t.Error("expected field error")
	}
}
