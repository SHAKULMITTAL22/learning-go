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

func Test___Type_possibleTypes_ebbe4d5f14(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	field := graphql.CollectedField{
		Name: "__Type",
		Type: introspection.Type{},
		Args: nil,
	}

	obj := &introspection.Type{}

	res := ec.___Type_possibleTypes(context.Background(), field, obj)

	if res == nil {
		t.Error("expected non-nil result")
	}
}

func Test___Type_possibleTypes_withError(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return nil, errors.New("some error")
		},
	}

	field := graphql.CollectedField{
		Name: "__Type",
		Type: introspection.Type{},
		Args: nil,
	}

	obj := &introspection.Type{}

	res := ec.___Type_possibleTypes(context.Background(), field, obj)

	if res != nil {
		t.Error("expected nil result")
	}
}
