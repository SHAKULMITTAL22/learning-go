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

func Test___EnumValue_name_d68367a5df(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	// Success case
	obj := &introspection.EnumValue{
		Name: "EnumValueName",
	}
	field := graphql.CollectedField{
		Selections: []*ast.Field{
			{
				Name: "name",
			},
		},
	}
	ctx := context.Background()
	res := ec.___EnumValue_name(ctx, field, obj)
	if res == nil {
		t.Error("Expected non-nil result")
	}
	if res.(string) != "EnumValueName" {
		t.Errorf("Expected \"EnumValueName\", got %s", res.(string))
	}

	// Error case
	obj = &introspection.EnumValue{
		Name: "",
	}
	res = ec.___EnumValue_name(ctx, field, obj)
	if res != nil {
		t.Error("Expected nil result")
	}
	if !graphql.HasFieldError(ctx, &graphql.FieldContext{
		Object:   "__EnumValue",
		Field:    field,
		Args:     nil,
		IsMethod: false,
	}) {
		t.Error("Expected field error")
	}
}
