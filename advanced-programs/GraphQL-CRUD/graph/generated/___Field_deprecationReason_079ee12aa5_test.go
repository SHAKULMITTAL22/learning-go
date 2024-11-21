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

func Test___Field_deprecationReason_079ee12aa5(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	field := graphql.CollectedField{
		Name: "__Field",
		Type: introspection.Field{},
		Args: nil,
	}

	obj := &introspection.Field{
		DeprecationReason: func() *string {
			s := "deprecated"
			return &s
		},
	}

	ret := ec.___Field_deprecationReason(context.Background(), field, obj)

	if ret == nil {
		t.Error("expected non-nil return value")
	}

	if *ret.(*string) != "deprecated" {
		t.Errorf("expected return value to be \"deprecated\", got %q", *ret.(*string))
	}
}

func Test___Field_deprecationReason_nil(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	field := graphql.CollectedField{
		Name: "__Field",
		Type: introspection.Field{},
		Args: nil,
	}

	obj := &introspection.Field{
		DeprecationReason: func() *string {
			return nil
		},
	}

	ret := ec.___Field_deprecationReason(context.Background(), field, obj)

	if ret != nil {
		t.Error("expected nil return value")
	}
}
