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

func Test___Type_inputFields_4325ed1a90(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	field := graphql.CollectedField{
		Name: "inputFields",
	}

	obj := &introspection.Type{
		Name: "__Type",
		InputFields: func() []introspection.InputValue {
			return []introspection.InputValue{
				{
					Name: "name",
					Type: &introspection.Type{
						Name: "String",
					},
				},
				{
					Name: "description",
					Type: &introspection.Type{
						Name: "String",
					},
				},
				{
					Name: "type",
					Type: &introspection.Type{
						Name: "String",
					},
				},
				{
					Name: "defaultValue",
					Type: &introspection.Type{
						Name: "String",
					},
				},
			}
		},
	}

	ret := ec.___Type_inputFields(context.Background(), field, obj)

	if _, ok := ret.(graphql.Marshaler); !ok {
		t.Error("return value is not a graphql.Marshaler")
	}

	// TODO: check the actual value of the return value
}

func Test___Type_inputFields_error(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	field := graphql.CollectedField{
		Name: "inputFields",
	}

	obj := &introspection.Type{
		Name: "__Type",
		InputFields: func() []introspection.InputValue {
			return nil
		},
	}

	ret := ec.___Type_inputFields(context.Background(), field, obj)

	if _, ok := ret.(graphql.Marshaler); !ok {
		t.Error("return value is not a graphql.Marshaler")
	}

	// TODO: check the actual value of the return value
}
