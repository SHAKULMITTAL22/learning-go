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

func Test_Item_id_4c0f4533a5(t *testing.T) {
	t.Run("Success Case", func(t *testing.T) {
		ec := &executionContext{}
		obj := &model.Item{ID: 1}
		field := graphql.CollectedField{}
		ctx := context.Background()

		result := ec._Item_id(ctx, field, obj)
		if result == graphql.Null {
			t.Errorf("Test Failed: Expected not null, received: %v", result)
		} else {
			t.Logf("Test Success: Expected not null, received: %v", result)
		}
	})

	t.Run("Failure Case - Recover from panic", func(t *testing.T) {
		ec := &executionContext{}
		obj := &model.Item{}
		field := graphql.CollectedField{}
		ctx := context.Background()

		result := ec._Item_id(ctx, field, obj)
		if result != graphql.Null {
			t.Errorf("Test Failed: Expected null, received: %v", result)
		} else {
			t.Logf("Test Success: Expected null, received: %v", result)
		}
	})

	t.Run("Failure Case - Error from middleware", func(t *testing.T) {
		ec := &executionContext{
			ResolverMiddleware: func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
				return nil, errors.New("middleware error")
			},
		}
		obj := &model.Item{ID: 1}
		field := graphql.CollectedField{}
		ctx := context.Background()

		result := ec._Item_id(ctx, field, obj)
		if result != graphql.Null {
			t.Errorf("Test Failed: Expected null, received: %v", result)
		} else {
			t.Logf("Test Success: Expected null, received: %v", result)
		}
	})
}
