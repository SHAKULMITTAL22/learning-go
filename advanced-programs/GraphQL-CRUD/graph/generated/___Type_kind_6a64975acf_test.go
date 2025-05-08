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

func Test___Type_kind_6a64975acf(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	// Success case
	obj := &introspection.Type{
		Kind: "OBJECT",
	}
	field := graphql.CollectedField{
		Selections: []*ast.Field{
			{
				Name: "kind",
			},
		},
	}
	ctx := context.Background()
	res := ec.___Type_kind(ctx, field, obj)
	if res == nil {
		t.Error("expected non-nil result")
	}
	if res.(string) != "OBJECT" {
		t.Errorf("expected kind to be OBJECT, got %s", res.(string))
	}

	// Error case
	obj = &introspection.Type{
		Kind: "",
	}
	res = ec.___Type_kind(ctx, field, obj)
	if res != nil {
		t.Error("expected nil result")
	}
	if !graphql.HasFieldError(ctx, &graphql.FieldContext{
		Object:   "__Type",
		Field:    field,
		Args:     nil,
		IsMethod: true,
	}) {
		t.Error("expected field error")
	}
}
