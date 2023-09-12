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

func Test___InputValue_type_793aa779f9(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	// Success case
	obj := &introspection.InputValue{
		Type: &introspection.Type{
			Kind: "SCALAR",
			Name: "String",
		},
	}
	field := graphql.CollectedField{
		Selections: []*ast.Field{
			{
				Name: "type",
			},
		},
	}
	ctx := context.Background()
	ret := ec.___InputValue_type(ctx, field, obj)
	if ret == nil {
		t.Error("expected non-nil return value")
	}

	// Error case
	obj = &introspection.InputValue{
		Type: nil,
	}
	ret = ec.___InputValue_type(ctx, field, obj)
	if ret != nil {
		t.Error("expected nil return value")
	}
}
