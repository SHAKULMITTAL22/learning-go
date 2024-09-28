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
	"github.com/99designs/gqlgen/graphql/introspection/testutil"
	"github.com/99designs/gqlgen/graphql/testutil"
)

func Test___Field_description_ee4af3c1b7(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		ec := &executionContext{}
		field := graphql.CollectedField{}
		obj := &introspection.Field{Description: "test"}

		result := ec.___Field_description(context.Background(), field, obj)
		if result == nil {
			t.Error("Expected non-nil result, got nil")
			t.Log("Method arguments: ", field, obj)
		} else {
			t.Log("Test success")
		}
	})

	t.Run("error case", func(t *testing.T) {
		ec := &executionContext{}
		field := graphql.CollectedField{}
		obj := &introspection.Field{Description: "test"}
		ec.ResolverMiddleware = func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
			return nil, errors.New("test error")
		}

		result := ec.___Field_description(context.Background(), field, obj)
		if result != graphql.Null {
			t.Error("Expected graphql.Null, got: ", result)
			t.Log("Method arguments: ", field, obj)
		} else {
			t.Log("Test success")
		}
	})

	t.Run("panic recovery", func(t *testing.T) {
		ec := &executionContext{}
		field := graphql.CollectedField{}
		obj := &introspection.Field{Description: "test"}
		ec.ResolverMiddleware = func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
			panic("test panic")
		}

		result := ec.___Field_description(context.Background(), field, obj)
		if result != graphql.Null {
			t.Error("Expected graphql.Null, got: ", result)
			t.Log("Method arguments: ", field, obj)
		} else {
			t.Log("Test success")
		}
	})
}
