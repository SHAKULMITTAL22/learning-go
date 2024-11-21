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

func Test___Directive_name_ca4c018cf8(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
	}

	field := graphql.CollectedField{
		Name: "name",
	}

	obj := &introspection.Directive{
		Name: "MyDirective",
	}

	ctx := context.Background()
	res := ec.___Directive_name(ctx, field, obj)

	if res == nil {
		t.Error("expected non-nil result")
	}

	if res.(string) != "MyDirective" {
		t.Errorf("expected \"MyDirective\", got %q", res.(string))
	}
}

func Test___Directive_name_error(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return nil, errors.New("some error")
		},
	}

	field := graphql.CollectedField{
		Name: "name",
	}

	obj := &introspection.Directive{
		Name: "MyDirective",
	}

	ctx := context.Background()
	res := ec.___Directive_name(ctx, field, obj)

	if res != nil {
		t.Error("expected nil result")
	}
}
