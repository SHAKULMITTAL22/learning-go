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

func Test___Field_name_bdd863490b(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	field := graphql.CollectedField{
		Name: "__Field",
		Selections: []*ast.Field{
			{
				Name: "name",
			},
		},
	}

	obj := &introspection.Field{
		Name: "TestField",
	}

	ctx := context.Background()
	res := ec.___Field_name(ctx, field, obj)

	if res == nil {
		t.Error("expected non-nil result")
	}

	if res.(string) != "TestField" {
		t.Errorf("expected result to be \"TestField\", got %s", res.(string))
	}
}

func Test___Field_name_with_error(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return nil, errors.New("some error")
		},
	}

	field := graphql.CollectedField{
		Name: "__Field",
		Selections: []*ast.Field{
			{
				Name: "name",
			},
		},
	}

	obj := &introspection.Field{
		Name: "TestField",
	}

	ctx := context.Background()
	res := ec.___Field_name(ctx, field, obj)

	if res != nil {
		t.Error("expected nil result")
	}
}
