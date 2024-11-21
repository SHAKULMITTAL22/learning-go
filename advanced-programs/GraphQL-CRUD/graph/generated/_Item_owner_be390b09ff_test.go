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

func Test_Item_owner_be390b09ff(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	// Success case
	obj := &model.Item{
		Owner: "John Doe",
	}
	field := graphql.CollectedField{
		Selections: []*ast.Field{
			{
				Name: "owner",
			},
		},
	}
	ctx := context.Background()
	ret := ec._Item_owner(ctx, field, obj)
	if ret == nil {
		t.Error("expected non-nil return value")
	}
	if ret.(string) != "John Doe" {
		t.Errorf("expected owner to be 'John Doe', got '%s'", ret.(string))
	}

	// Error case
	obj = &model.Item{
		Owner: "",
	}
	ret = ec._Item_owner(ctx, field, obj)
	if ret != nil {
		t.Error("expected nil return value")
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
