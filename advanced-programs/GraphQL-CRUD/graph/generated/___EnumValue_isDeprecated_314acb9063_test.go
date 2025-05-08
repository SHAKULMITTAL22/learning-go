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

func Test___EnumValue_isDeprecated_314acb9063(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	// Success case
	obj := &introspection.EnumValue{
		Name: "EnumValue",
		DeprecationReason: "Deprecated",
	}
	field := graphql.CollectedField{
		Selections: []*ast.Field{
			{
				Name: "isDeprecated",
			},
		},
	}
	ctx := context.Background()
	res := ec.___EnumValue_isDeprecated(ctx, field, obj)
	if res == nil {
		t.Error("expected non-nil result")
	}
	if !res.(bool) {
		t.Error("expected true, got false")
	}

	// Failure case
	obj = &introspection.EnumValue{
		Name: "EnumValue",
	}
	res = ec.___EnumValue_isDeprecated(ctx, field, obj)
	if res != nil {
		t.Error("expected nil result")
	}
}
