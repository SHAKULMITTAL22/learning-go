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

func Test___Schema_mutationType_28bfa8c8a6(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	// Success case
	obj := &introspection.Schema{
		MutationType: func() *introspection.Type {
			return &introspection.Type{
				Name: "Mutation",
			}
		},
	}
	field := graphql.CollectedField{
		Selections: []*ast.Field{
			{
				Name: "mutationType",
			},
		},
	}
	ctx := context.Background()
	res := ec.___Schema_mutationType(ctx, field, obj)
	if res == nil {
		t.Error("expected non-nil result")
	}

	// Error case
	obj = &introspection.Schema{
		MutationType: func() *introspection.Type {
			return nil
		},
	}
	res = ec.___Schema_mutationType(ctx, field, obj)
	if res != nil {
		t.Error("expected nil result")
	}
}
