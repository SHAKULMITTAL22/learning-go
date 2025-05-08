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

func Test_Query_item_f90ef048de(t *testing.T) {
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
		Variables: make(graphql.Variables),
	}

	query := `
	query {
		item(id: 1) {
			id
			name
			description
		}
	}
	`

	rawArgs := map[string]interface{}{
		"id": 1,
	}

	args, err := ec.field_Query_item_args(context.Background(), rawArgs)
	if err != nil {
		t.Error(err)
		return
	}

	fc := &graphql.FieldContext{
		Object:   "Query",
		Field:    graphql.CollectedField{},
		Args:     args,
		IsMethod: true,
	}

	ctx := graphql.WithFieldContext(context.Background(), fc)

	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		return ec.resolvers.Query().Item(rctx, args["id"].(int))
	})
	if err != nil {
		t.Error(err)
		return
	}

	res := resTmp.(*model.Item)

	if res == nil {
		t.Error("expected non-nil result")
		return
	}

	if res.ID != 1 {
		t.Errorf("expected ID to be 1, got %d", res.ID)
	}

	if res.Name != "Item 1" {
		t.Errorf("expected Name to be \"Item 1\", got \"%s\"", res.Name)
	}

	if res.Description != "This is item 1" {
		t.Errorf("expected Description to be \"This is item 1\", got \"%s\"", res.Description)
	}
}

func Test_Query_item_error(t *testing.T) {
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
		Variables: make(graphql.Variables),
	}

	query := `
	query {
		item(id: 1) {
			id
			name
			description
		}
	}
	`

	rawArgs := map[string]interface{}{
		"id": "not an int",
	}

	args, err := ec.field_Query_item_args(context.Background(), rawArgs)
	if err != nil {
		t.Error(err)
		return
	}

	fc := &graphql.FieldContext{
		Object:   "Query",
		Field:    graphql.CollectedField{},
		Args:     args,
		IsMethod: true,
	}

	ctx := graphql.WithFieldContext(context.Background(), fc)

	_, err = ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		return ec.resolvers.Query().Item(rctx, args["id"].(int))
	})
	if err == nil {
		t.Error("expected error, got nil")
		return
	}

	if !errors.Is(err, strconv.ErrSyntax) {
		t.Errorf("expected error to be strconv.ErrSyntax, got %v", err)
	}
}
