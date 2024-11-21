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

func Test___Type_name_0e76cd4719(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	obj := &introspection.Type{
		Name: "TestType",
	}

	field := graphql.CollectedField{
		Selections: []*ast.Field{
			{
				Name: "name",
			},
		},
	}

	ctx := context.Background()
	res := ec.___Type_name(ctx, field, obj)

	if res == nil {
		t.Error("Expected a non-nil result")
	}

	if res.(string) != "TestType" {
		t.Errorf("Expected the name to be 'TestType', but got '%s'", res.(string))
	}
}

func Test___Type_name_with_error(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return nil, errors.New("some error")
		},
	}

	obj := &introspection.Type{
		Name: "TestType",
	}

	field := graphql.CollectedField{
		Selections: []*ast.Field{
			{
				Name: "name",
			},
		},
	}

	ctx := context.Background()
	res := ec.___Type_name(ctx, field, obj)

	if res != nil {
		t.Error("Expected a nil result")
	}
}
