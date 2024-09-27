package generated

import (
	"context"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/graph/model"
)

type MutationResolver struct {
	CreateItems func(ctx context.Context, input model.NewItem) (*model.Item, error)
	results     *model.Item
}

func (_ *MutationResolver) Mutation() MutationResolver {
	return MutationResolver{}
}

type executionContext struct {
	ResolverMiddleware func(ctx context.Context, next graphql.Resolver) (res interface{}, err error)
	resolvers          *MutationResolver
}

func (ec *executionContext) _Mutation_createItems(ctx context.Context, field graphql.CollectedField) (response interface{}) {
	next := func(ctx context.Context) (interface{}, error) {
		return ec.resolvers.CreateItems(ctx, model.NewItem{})
	}
	res, err := ec.ResolverMiddleware(ctx, next)
	if err != nil || res == nil {
		return graphql.Null
	}

	return res
}

func Test_Mutation_createItems_f1177e58af(t *testing.T) {
	t.Run("it should return error when model.NewItem is nil", func(t *testing.T) {
		ctx := context.Background()
		field := graphql.CollectedField{}
		mockResolver := &MutationResolver{}

		ec := &executionContext{
			ResolverMiddleware: func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
				return next(ctx)
			},
			resolvers: mockResolver,
		}

		response := ec._Mutation_createItems(ctx, field)
		assert.Equal(t, graphql.Null, response, "should return graphql.Null for nil argument")
	})

	t.Run("it should return result when NewItem is correct", func(t *testing.T) {
		ctx := context.Background()
		field := graphql.CollectedField{}

		resTemp := &model.Item{
			ID:          "1",
			Title:       "Test title",
			Description: "Test description",
		}

		mockResolver := &MutationResolver{
			CreateItems: func(ctx context.Context, input model.NewItem) (*model.Item, error) {
				return resTemp, nil
			},
		}

		ec := &executionContext{
			ResolverMiddleware: func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
				return next(ctx)
			},
			resolvers: mockResolver,
		}

		response := ec._Mutation_createItems(ctx, field)
		assert.Equal(t, resTemp, response, "should return expected result")
	})
}
