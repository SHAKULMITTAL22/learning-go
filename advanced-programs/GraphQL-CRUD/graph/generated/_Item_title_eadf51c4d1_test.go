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

func Test_Item_title_eadf51c4d1(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	// Success case
	obj := &model.Item{
		Title: "Test Title",
	}
	field := graphql.CollectedField{
		Selections: []*ast.Field{
			{
				Name: "title",
			},
		},
	}
	ctx := context.Background()
	res := ec._Item_title(ctx, field, obj)
	if res == nil {
		t.Error("Expected non-nil result")
	}
	if res.(string) != "Test Title" {
		t.Errorf("Expected title to be 'Test Title', got '%s'", res.(string))
	}

	// Error case
	obj = &model.Item{}
	res = ec._Item_title(ctx, field, obj)
	if res != nil {
		t.Error("Expected nil result")
	}
	if !graphql.HasFieldError(ctx, &graphql.FieldContext{
		Object:   "Item",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}) {
		t.Error("Expected field error")
	}
}
