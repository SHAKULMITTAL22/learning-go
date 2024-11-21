package generated

import (
	"bytes"
	"context"
	"errors"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/graph/model"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

func Test_Item_owner_be390b09ff(t *testing.T) {
	t.Run("Success Case", func(t *testing.T) {
		ec := &executionContext{}
		item := &model.Item{Owner: "John Doe"}
		field := graphql.CollectedField{}
		ctx := context.Background()
		result := ec._Item_owner(ctx, field, item)

		if result == nil {
			t.Error("Expected non-nil value, got nil")
			t.Log("Method arguments: ", ctx, field, item)
		} else {
			t.Log("Test Success")
		}
	})

	t.Run("Failure Case - Recover from Panic", func(t *testing.T) {
		ec := &executionContext{}
		item := &model.Item{Owner: "John Doe"}
		field := graphql.CollectedField{}
		ctx := context.Background()

		defer func() {
			if r := recover(); r != nil {
				t.Log("Recovered from panic")
			} else {
				t.Error("Expected a panic, but did not occur")
				t.Log("Method arguments: ", ctx, field, item)
			}
		}()

		ec._Item_owner(ctx, field, item)
	})

	t.Run("Failure Case - Error in ResolverMiddleware", func(t *testing.T) {
		ec := &executionContext{
			ResolverMiddleware: func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
				return nil, errors.New("Test error")
			},
		}
		item := &model.Item{Owner: "John Doe"}
		field := graphql.CollectedField{}
		ctx := context.Background()

		result := ec._Item_owner(ctx, field, item)

		if result != nil {
			t.Error("Expected nil value due to error, got non-nil")
			t.Log("Method arguments: ", ctx, field, item)
		} else {
			t.Log("Test Success")
		}
	})
}
