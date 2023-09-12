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

func Test_Item_rating_92861babf7(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	// Success case
	obj := &model.Item{
		Rating: 5,
	}
	field := graphql.CollectedField{
		Selections: []*ast.Field{
			{
				Name: "rating",
			},
		},
	}
	ctx := context.Background()
	res := ec._Item_rating(ctx, field, obj)
	if res == nil {
		t.Error("expected non-nil result")
	}
	if res.(int) != 5 {
		t.Errorf("expected rating to be 5, got %d", res.(int))
	}

	// Error case
	obj = &model.Item{}
	res = ec._Item_rating(ctx, field, obj)
	if res != nil {
		t.Error("expected nil result")
	}
	if !graphql.HasFieldError(ctx, &graphql.FieldContext{
		Object:   "Item",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}) {
		t.Error("expected field error")
	}
}
