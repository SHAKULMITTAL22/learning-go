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
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

type executionContextTest struct {
	ResolverMiddleware func(ctx context.Context, next graphql.Resolver) (res interface{}, err error)
}

func (ec *executionContextTest) ___Field_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	return nil
}

func Test___Field_name_bdd863490b(t *testing.T) {
	ec := &executionContextTest{
		ResolverMiddleware: func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
			return "test", nil
		},
	}

	field := graphql.CollectedField{}
	obj := &introspection.Field{
		Name: "test",
	}

	t.Run("success", func(t *testing.T) {
		ret := ec.___Field_name(context.Background(), field, obj)
		if ret == nil || ret.MarshalJSON() != "\"test\"" {
			t.Errorf("Failed: got %v, expected \"test\"", ret)
		} else {
			t.Log("Success: Expected result obtained")
		}
	})

	t.Run("error", func(t *testing.T) {
		ec.ResolverMiddleware = func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
			return nil, errors.New("test error")
		}

		ret := ec.___Field_name(context.Background(), field, obj)
		if ret != nil {
			t.Errorf("Failed: got %v, expected nil", ret)
		} else {
			t.Log("Success: Expected result obtained")
		}
	})
}
