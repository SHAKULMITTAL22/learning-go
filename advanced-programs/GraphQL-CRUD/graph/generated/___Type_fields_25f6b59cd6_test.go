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

func Test___Type_fields_25f6b59cd6(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	// Success case
	field := graphql.CollectedField{
		Name: "__Type",
		Type: introspection.Type{},
		Args: map[string]interface{}{
			"includeDeprecated": true,
		},
	}
	obj := &introspection.Type{}
	res := ec.___Type_fields(context.Background(), field, obj)
	if res == nil {
		t.Error("expected non-nil result")
	}

	// Error case
	field = graphql.CollectedField{
		Name: "__Type",
		Type: introspection.Type{},
		Args: map[string]interface{}{
			"includeDeprecated": "not a bool",
		},
	}
	obj = &introspection.Type{}
	res = ec.___Type_fields(context.Background(), field, obj)
	if res != nil {
		t.Error("expected nil result")
	}
}
