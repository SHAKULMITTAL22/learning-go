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
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"

	"testing"
)

// Unit test for the executionContext.___EnumValue_isDeprecated method
func Test___EnumValue_isDeprecated_314acb9063(t *testing.T) {
	t.Run("Test with valid input", func(t *testing.T) {
		ec := newExecutionContext()
		ctx := context.Background()
		field := graphql.CollectedField{}
		obj := &introspection.EnumValue{IsDeprecated: true}

		ret := ec.___EnumValue_isDeprecated(ctx, field, obj)
		if ret == graphql.Null {
			t.Error("Expected boolean, but got null")
		} else {
			t.Log("Test with valid input passed")
		}
	})

	t.Run("Test with null return from ResolverMiddleware", func(t *testing.T) {
		ec := newExecutionContext()
		ec.ResolverMiddleware = func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
			return nil, nil
		}
		ctx := context.Background()
		field := graphql.CollectedField{}
		obj := &introspection.EnumValue{IsDeprecated: true}

		ret := ec.___EnumValue_isDeprecated(ctx, field, obj)
		if ret != graphql.Null {
			t.Error("Expected null, but got boolean")
		} else {
			t.Log("Test with null return from ResolverMiddleware passed")
		}
	})

	t.Run("Test with error return from ResolverMiddleware", func(t *testing.T) {
		ec := newExecutionContext()
		ec.ResolverMiddleware = func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
			return nil, errors.New("Test error")
		}
		ctx := context.Background()
		field := graphql.CollectedField{}
		obj := &introspection.EnumValue{IsDeprecated: true}

		ret := ec.___EnumValue_isDeprecated(ctx, field, obj)
		if ret != graphql.Null {
			t.Error("Expected null, but got boolean")
		} else {
			t.Log("Test with error return from ResolverMiddleware passed")
		}
	})
}
