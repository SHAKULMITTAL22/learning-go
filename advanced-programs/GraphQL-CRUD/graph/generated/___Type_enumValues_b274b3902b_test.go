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

func Test___Type_enumValues_b274b3902b(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	// Success case
	obj := &introspection.Type{
		Name: "MyType",
		EnumValues: []introspection.EnumValue{
			{
				Name: "EnumValue1",
				Description: "Description for EnumValue1",
			},
			{
				Name: "EnumValue2",
				Description: "Description for EnumValue2",
			},
		},
	}
	field := graphql.CollectedField{
		Name: "enumValues",
		ArgumentMap: map[string]interface{}{
			"includeDeprecated": false,
		},
	}
	ctx := context.Background()
	res := ec.___Type_enumValues(ctx, field, obj)
	if res == nil {
		t.Error("Expected non-nil result")
	}
	if len(res.([]introspection.EnumValue)) != 2 {
		t.Errorf("Expected 2 enum values, got %d", len(res.([]introspection.EnumValue)))
	}

	// Error case
	obj = &introspection.Type{
		Name: "MyType",
		EnumValues: nil,
	}
	res = ec.___Type_enumValues(ctx, field, obj)
	if res != nil {
		t.Error("Expected nil result")
	}
}
