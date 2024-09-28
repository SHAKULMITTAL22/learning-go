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

func Test___Directive_description_d107698109(t *testing.T) {
	ec := &executionContext{}
	field := graphql.CollectedField{}
	ctx := context.Background()

	t.Run("Test with valid directive", func(t *testing.T) {
		obj := &introspection.Directive{
			Description: "Test directive",
		}
		result := ec.___Directive_description(ctx, field, obj)
		if result == graphql.Null {
			t.Error("Test with valid directive failed")
		} else {
			t.Log("Test with valid directive passed")
		}
	})

	t.Run("Test with nil directive", func(t *testing.T) {
		obj := &introspection.Directive{
			Description: nil,
		}
		result := ec.___Directive_description(ctx, field, obj)
		if result != graphql.Null {
			t.Error("Test with nil directive failed")
		} else {
			t.Log("Test with nil directive passed")
		}
	})

	t.Run("Test with panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Error("Test with panic failed")
			} else {
				t.Log("Test with panic passed")
			}
		}()
		obj := &introspection.Directive{
			Description: "Test directive",
		}
		ec.ResolverMiddleware = func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
			panic("Test panic")
		}
		_ = ec.___Directive_description(ctx, field, obj)
	})
}
