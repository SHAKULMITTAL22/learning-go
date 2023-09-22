package generated

import (
	"context"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/vektah/gqlparser/v2/ast"
)

type executionContext struct{}

func (ec *executionContext) ___Type(ctx context.Context, sel ast.SelectionSet, v *introspection.Type) graphql.Marshaler {
	return graphql.Null
}

func (ec *executionContext) marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v *introspection.Type) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, sel, v)
}

func TestType_6b1e5496f4(t *testing.T) {
	ctx := context.Background()
	sel := ast.SelectionSet{}
	ec := &executionContext{}

	t.Run("When v is nil", func(t *testing.T) {
		v := (*introspection.Type)(nil)
		res := ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, sel, v)
		if res != graphql.Null {
			t.Errorf("Expected graphql.Null, got %v", res)
		}
	})

	t.Run("When v is not nil", func(t *testing.T) {
		v := &introspection.Type{}
		res := ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, sel, v)
		if res == graphql.Null {
			t.Errorf("Expected not to be graphql.Null, but got %v", res)
		}
	})
}
