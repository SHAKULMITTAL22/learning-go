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

func Test___Directive_description_d107698109(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	field := graphql.CollectedField{
		Name: "__Directive",
		Type: introspection.Directive{},
		Args: nil,
	}

	obj := &introspection.Directive{
		Description: "A directive",
	}

	ret := ec.___Directive_description(context.Background(), field, obj)

	if ret == nil {
		t.Error("expected non-nil return value")
	}

	if ret.(string) != "A directive" {
		t.Errorf("expected return value to be \"A directive\", got %q", ret.(string))
	}
}

func Test___Directive_description_with_error(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return nil, errors.New("some error")
		},
	}

	field := graphql.CollectedField{
		Name: "__Directive",
		Type: introspection.Directive{},
		Args: nil,
	}

	obj := &introspection.Directive{
		Description: "A directive",
	}

	ret := ec.___Directive_description(context.Background(), field, obj)

	if ret != nil {
		t.Error("expected nil return value")
	}
}
