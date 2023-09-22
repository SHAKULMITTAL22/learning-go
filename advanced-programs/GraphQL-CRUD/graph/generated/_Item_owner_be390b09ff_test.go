package generated

import (
	"context"
	"errors"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/graph/model"
)

type executionContext struct {
	ResolverMiddleware func(ctx context.Context, next graphql.Resolver) (res interface{}, err error)
}

func newExecutionContext(resolverMiddleware func(ctx context.Context, next graphql.Resolver) (res interface{}, err error)) *executionContext {
	return &executionContext{
		ResolverMiddleware: resolverMiddleware,
	}
}

func (ec *executionContext) _Item_owner(ctx context.Context, field graphql.CollectedField, obj *model.Item) (res interface{}) {
	if obj.Owner == nil {
		return nil
	}
	return obj.Owner
}

func Test_Item_owner_be390b09ff(t *testing.T) {
	ec := newExecutionContext(nil)

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		field := graphql.CollectedField{}
		obj := &model.Item{Owner: "John"}

		res := ec._Item_owner(ctx, field, obj)

		if res == nil {
			t.Errorf("Expected not null, got %v", res)
		}
		t.Log("Test_Item_owner_be390b09ff success case passed")
	})

	t.Run("error", func(t *testing.T) {
		ec := newExecutionContext(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
			return nil, errors.New("resolver error")
		})

		ctx := context.Background()
		field := graphql.CollectedField{}
		obj := &model.Item{Owner: "John"}

		res := ec._Item_owner(ctx, field, obj)

		if res != nil {
			t.Errorf("Expected nil, got %v", res)
		}
		t.Log("Test_Item_owner_be390b09ff error case passed")
	})

	t.Run("nil model", func(t *testing.T) {
		ctx := context.Background()
		field := graphql.CollectedField{}
		obj := &model.Item{Owner: nil}

		res := ec._Item_owner(ctx, field, obj)

		if res != nil {
			t.Errorf("Expected nil, got %v", res)
		}
		t.Log("Test_Item_owner_be390b09ff passed with nil model")
	})
}
