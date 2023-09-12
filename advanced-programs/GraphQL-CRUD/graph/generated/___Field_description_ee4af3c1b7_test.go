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

func Test___Field_description_ee4af3c1b7(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	field := graphql.CollectedField{
		Name: "__Field",
		Type: introspection.Field{},
	}

	obj := &introspection.Field{
		Description: "description",
	}

	ctx := context.Background()
	res := ec.___Field_description(ctx, field, obj)

	if res == nil {
		t.Error("expected non-nil result")
	}

	if res.(string) != "description" {
		t.Errorf("expected description to be 'description', got '%s'", res.(string))
	}
}

func Test___Field_description_with_error(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(context.Context) (interface{}, error)) (interface{}, error) {
			return nil, errors.New("error")
		},
	}

	field := graphql.CollectedField{
		Name: "__Field",
		Type: introspection.Field{},
	}

	obj := &introspection.Field{
		Description: "description",
	}

	ctx := context.Background()
	res := ec.___Field_description(ctx, field, obj)

	if res != nil {
		t.Error("expected nil result")
	}
}
