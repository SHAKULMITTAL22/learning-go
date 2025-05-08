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

func Test___EnumValue_deprecationReason_75526065d4(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	obj := &introspection.EnumValue{
		DeprecationReason: func() *string {
			s := "deprecation reason"
			return &s
		},
	}

	field := graphql.CollectedField{
		Selections: []*ast.Field{
			{
				Name: "deprecationReason",
			},
		},
	}

	ctx := context.Background()
	res := ec.___EnumValue_deprecationReason(ctx, field, obj)

	if res == nil {
		t.Error("expected non-nil result")
	}

	s, ok := res.(string)
	if !ok {
		t.Errorf("expected string result, got %T", res)
	}

	if s != "deprecation reason" {
		t.Errorf("expected deprecation reason to be 'deprecation reason', got '%s'", s)
	}
}

func Test___EnumValue_deprecationReason_nil(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	obj := &introspection.EnumValue{
		DeprecationReason: func() *string {
			return nil
		},
	}

	field := graphql.CollectedField{
		Selections: []*ast.Field{
			{
				Name: "deprecationReason",
			},
		},
	}

	ctx := context.Background()
	res := ec.___EnumValue_deprecationReason(ctx, field, obj)

	if res != nil {
		t.Error("expected nil result")
	}
}
