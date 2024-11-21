package generated

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"testing"
)

type mockResolverMiddleware struct {
	result *introspection.Type
	err    error
}

func (m *mockResolverMiddleware) Middleware(next graphql.Resolver) graphql.Resolver {
	return func(ctx context.Context) (res interface{}, err error) {
		return m.result, m.err
	}
}

func Test_Query___type_5847112bcc(t *testing.T) {
	t.Run("Test with valid arguments", func(t *testing.T) {
		ec := &executionContext{
			ResolverMiddleware: &mockResolverMiddleware{
				result: &introspection.Type{Name: "test"},
			},
		}
		field := graphql.CollectedField{}
		res := ec._Query___type(context.Background(), field)
		if res == graphql.Null {
			t.Error("Expected non null response")
		}
		t.Log("Test with valid arguments passed")
	})

	t.Run("Test with invalid arguments", func(t *testing.T) {
		ec := &executionContext{
			ResolverMiddleware: &mockResolverMiddleware{
				err: errors.New("error"),
			},
		}
		field := graphql.CollectedField{}
		res := ec._Query___type(context.Background(), field)
		if res != graphql.Null {
			t.Error("Expected null response")
		}
		t.Log("Test with invalid arguments passed")
	})
}
