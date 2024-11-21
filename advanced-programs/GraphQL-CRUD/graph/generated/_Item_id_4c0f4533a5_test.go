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

func Test_Item_id_4c0f4533a5(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	obj := &model.Item{
		ID: 1,
	}

	field := graphql.CollectedField{
		Selections: []*ast.Field{
			{
				Name: "id",
			},
		},
	}

	ctx := context.Background()
	res := ec._Item_id(ctx, field, obj)

	if res == nil {
		t.Error("expected non-nil result")
	}

	if res.(int) != 1 {
		t.Errorf("expected id to be 1, got %d", res.(int))
	}
}

func Test_Item_id_error(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return nil, errors.New("some error")
		},
	}

	obj := &model.Item{
		ID: 1,
	}

	field := graphql.CollectedField{
		Selections: []*ast.Field{
			{
				Name: "id",
			},
		},
	}

	ctx := context.Background()
	res := ec._Item_id(ctx, field, obj)

	if res != nil {
		t.Error("expected nil result")
	}
}
