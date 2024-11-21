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

func Test___Type_description_4bf5e43c80(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	// Success case
	obj := &introspection.Type{
		Description: "A type",
	}
	field := graphql.CollectedField{
		Selections: []*ast.Field{
			{
				Name: "description",
			},
		},
	}
	ctx := context.Background()
	res := ec.___Type_description(ctx, field, obj)
	if res == nil {
		t.Error("Expected non-nil result")
	}
	if res.(string) != "A type" {
		t.Errorf("Expected description to be 'A type', got '%s'", res.(string))
	}

	// Error case
	obj = &introspection.Type{
		Description: "",
	}
	res = ec.___Type_description(ctx, field, obj)
	if res != nil {
		t.Error("Expected nil result")
	}
}
