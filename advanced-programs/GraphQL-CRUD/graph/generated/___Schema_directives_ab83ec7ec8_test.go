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

func Test___Schema_directives_ab83ec7ec8(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	// Success case
	obj := &introspection.Schema{}
	obj.Directives = func() []introspection.Directive {
		return []introspection.Directive{
			{
				Name: "directive1",
				Args: []introspection.InputValue{
					{
						Name: "arg1",
						Type: "String",
						DefaultValue: "value1",
					},
				},
			},
			{
				Name: "directive2",
				Args: []introspection.InputValue{
					{
						Name: "arg2",
						Type: "Int",
						DefaultValue: 123,
					},
				},
			},
		}
	}
	res := ec.___Schema_directives(context.Background(), graphql.CollectedField{}, obj)
	if res == nil {
		t.Error("expected non-nil result")
	}
	if len(res.([]introspection.Directive)) != 2 {
		t.Errorf("expected 2 directives, got %d", len(res.([]introspection.Directive)))
	}

	// Error case
	obj.Directives = func() []introspection.Directive {
		return nil
	}
	res = ec.___Schema_directives(context.Background(), graphql.CollectedField{}, obj)
	if res != nil {
		t.Error("expected nil result")
	}
}
