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

func Test___Type_interfaces_cde63e4dc2(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	// Success case
	obj := &introspection.Type{
		Interfaces: func() []introspection.Type {
			return []introspection.Type{
				{
					Name: "Foo",
				},
				{
					Name: "Bar",
				},
			}
		},
	}
	field := graphql.CollectedField{
		Selections: []*ast.Field{
			{
				Name: "name",
			},
		},
	}
	ctx := context.Background()
	res := ec.___Type_interfaces(ctx, field, obj)
	if res == nil {
		t.Error("expected non-nil result")
	}
	if len(res.([]introspection.Type)) != 2 {
		t.Errorf("expected 2 interfaces, got %d", len(res.([]introspection.Type)))
	}

	// Error case
	obj = &introspection.Type{
		Interfaces: func() []introspection.Type {
			return nil
		},
	}
	res = ec.___Type_interfaces(ctx, field, obj)
	if res != nil {
		t.Error("expected nil result")
	}
}
