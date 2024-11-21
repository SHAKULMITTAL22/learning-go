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

func Test___Directive_locations_2d6c7100e1(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	field := graphql.CollectedField{
		Name: "__Directive",
		Type: introspection.Directive{},
	}

	obj := &introspection.Directive{
		Locations: []string{"QUERY"},
	}

	ret := ec.___Directive_locations(context.Background(), field, obj)

	if ret == nil {
		t.Error("expected non-nil return value")
	}

	if len(ret.([]string)) != 1 {
		t.Errorf("expected 1 location, got %d", len(ret.([]string)))
	}

	if ret.([]string)[0] != "QUERY" {
		t.Errorf("expected location to be QUERY, got %s", ret.([]string)[0])
	}
}

func Test___Directive_locations_error(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	field := graphql.CollectedField{
		Name: "__Directive",
		Type: introspection.Directive{},
	}

	obj := &introspection.Directive{
		Locations: nil,
	}

	ret := ec.___Directive_locations(context.Background(), field, obj)

	if ret != nil {
		t.Error("expected nil return value")
	}

	if !graphql.HasFieldError(context.Background(), &graphql.FieldContext{
		Object:   "__Directive",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}) {
		t.Error("expected field error")
	}
}
