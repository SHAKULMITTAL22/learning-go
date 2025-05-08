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

func Test_Query_items_9c4bcb9479(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
		Recover: func(ctx context.Context, err interface{}) error {
			return err
		},
		Error: func(ctx context.Context, err error) {
			t.Error(err)
		},
		Errorf: func(ctx context.Context, format string, args ...interface{}) {
			t.Errorf(format, args...)
		},
	}

	ctx := context.Background()
	field := graphql.CollectedField{
		Object:   "Query",
		Field:    ast.Field{Name: "items"},
		Args:     nil,
		IsMethod: true,
	}

	// Success case
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		return []*model.Item{
			{ID: 1, Name: "Item 1"},
			{ID: 2, Name: "Item 2"},
		}, nil
	})
	if err != nil {
		t.Error(err)
	}
	res := resTmp.([]*model.Item)
	if len(res) != 2 {
		t.Errorf("Expected 2 items, got %d", len(res))
	}

	// Error case
	resTmp, err = ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		return nil, errors.New("some error")
	})
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if resTmp != nil {
		t.Errorf("Expected nil, got %v", resTmp)
	}
}
