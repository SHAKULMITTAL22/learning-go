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

func Test_Mutation_delete_bae6487f43(t *testing.T) {
	// TODO: Change the value of id to a valid value
	id := 1
	// TODO: Change the value of expected to a valid value
	expected := "string"

	t.Run("success", func(t *testing.T) {
		ec := executionContext{
			Resolvers: &mockResolvers{
				Mutation: &mockMutationResolver{
					Delete: func(ctx context.Context, id int) (string, error) {
						return expected, nil
					},
				},
			},
		}

		rawArgs := map[string]interface{}{
			"id": id,
		}
		args, err := ec.field_Mutation_delete_args(context.Background(), rawArgs)
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

		ctx := graphql.WithFieldContext(context.Background(), fc)
		resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
			return ec.resolvers.Mutation().Delete(rctx, args["id"].(int))
		})
		if err != nil {
			t.Error(err)
			return
		}

		res := resTmp.(string)
		if res != expected {
			t.Errorf("expected %v, got %v", expected, res)
		}
	})

	t.Run("error", func(t *testing.T) {
		ec := executionContext{
			Resolvers: &mockResolvers{
				Mutation: &mockMutationResolver{
					Delete: func(ctx context.Context, id int) (string, error) {
						return "", errors.New("error")
					},
				},
			},
		}

		rawArgs := map[string]interface{}{
			"id": id,
		}
		args, err := ec.field_Mutation_delete_args(context.Background(), rawArgs)
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

		ctx := graphql.WithFieldContext(context.Background(), fc)
		resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
			return ec.resolvers.Mutation().Delete(rctx, args["id"].(int))
		})
		if err == nil {
			t.Error("expected error, got nil")
			return
		}

		if resTmp != nil {
			t.Errorf("expected nil, got %v", resTmp)
		}
	})
}
