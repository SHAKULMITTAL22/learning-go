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

func Test_Mutation_createItems_f1177e58af(t *testing.T) {
	ec := executionContext{
		ResolverMiddleware: func(ctx context.Context, next func(ctx context.Context) (interface{}, error)) (interface{}, error) {
			return next(ctx)
		},
		Recover: func(ctx context.Context, err interface{}) error {
			return err
		},
		Error: func(ctx context.Context, err error) {
			t.Error(err)
		},
		Variables: make(graphql.Variables),
	}

	ctx := context.Background()
	rawArgs := make(map[string]interface{})
	rawArgs["input"] = model.NewItem{
		Name:  "Test Item",
		Price: 100,
	}
	args, err := ec.field_Mutation_createItems_args(ctx, rawArgs)
	if err != nil {
		t.Error(err)
		return
	}

	fc := &graphql.FieldContext{
		Object:   "Mutation",
		Field:    graphql.CollectedField{},
		Args:     args,
		IsMethod: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		return ec.resolvers.Mutation().CreateItems(rctx, args["input"].(model.NewItem))
	})
	if err != nil {
		t.Error(err)
		return
	}
	if resTmp == nil {
		t.Error("must not be null")
		return
	}
	res := resTmp.(*model.Item)
	fc.Result = res

	// TODO: Check the result of the mutation

}
