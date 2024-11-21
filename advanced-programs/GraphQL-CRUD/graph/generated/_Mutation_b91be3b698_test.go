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
	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/graph/model"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

func Test_Mutation_b91be3b698(t *testing.T) {
	ec := executionContext{
		OperationContext: &graphql.OperationContext{
			Variables: map[string]interface{}{},
		},
	}

	sel := ast.SelectionSet{
		Selections: []ast.Selection{
			&ast.Field{
				Name: "createItems",
			},
		},
	}

	out := ec._Mutation(context.Background(), sel)
	if out == graphql.Null {
		t.Error("expected non-null result")
	}

	// TODO: assert the result is correct
}

func Test_Mutation_b91be3b698_with_error(t *testing.T) {
	ec := executionContext{
		OperationContext: &graphql.OperationContext{
			Variables: map[string]interface{}{},
		},
	}

	sel := ast.SelectionSet{
		Selections: []ast.Selection{
			&ast.Field{
				Name: "createItems",
			},
		},
	}

	ec.ResolverMiddleware = func(ctx context.Context, next graphql.Resolver) (res graphql.Marshaler) {
		return graphql.Error("some error")
	}

	out := ec._Mutation(context.Background(), sel)
	if out != graphql.Null {
		t.Error("expected null result")
	}
}
