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

func Test___InputValue_defaultValue_8d4dfa9f8b(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	// Success case
	obj := &introspection.InputValue{
		DefaultValue: "Hello World",
	}
	field := graphql.CollectedField{
		Selections: []*ast.Field{
			{
				Name: "defaultValue",
			},
		},
	}
	ctx := context.Background()
	ret := ec.___InputValue_defaultValue(ctx, field, obj)
	if ret == nil {
		t.Error("Expected non-nil return value")
	}
	if ret.(string) != "Hello World" {
		t.Errorf("Expected return value to be \"Hello World\", got %s", ret.(string))
	}

	// Error case
	obj = &introspection.InputValue{
		DefaultValue: nil,
	}
	ret = ec.___InputValue_defaultValue(ctx, field, obj)
	if ret != nil {
		t.Error("Expected nil return value")
	}
}
