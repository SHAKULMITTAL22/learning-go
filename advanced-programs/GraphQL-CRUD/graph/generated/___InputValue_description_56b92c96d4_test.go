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

func Test___InputValue_description_56b92c96d4(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	// Success case
	obj := &introspection.InputValue{
		Description: "This is a description",
	}
	field := graphql.CollectedField{
		Selections: []*ast.Field{
			{
				Name: "description",
			},
		},
	}
	ctx := context.Background()
	ret := ec.___InputValue_description(ctx, field, obj)
	if ret == nil {
		t.Error("Expected non-nil return value")
	}
	if ret.(string) != "This is a description" {
		t.Errorf("Expected description to be \"This is a description\", got %s", ret.(string))
	}

	// Error case
	obj = &introspection.InputValue{
		Description: "",
	}
	ret = ec.___InputValue_description(ctx, field, obj)
	if ret != nil {
		t.Error("Expected nil return value")
	}
}
