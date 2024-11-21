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

func Test_Query_5fade268b0(t *testing.T) {
	ec := executionContext{
		OperationContext: &graphql.OperationContext{
			Variables: map[string]interface{}{},
		},
	}

	sel := ast.SelectionSet{
		Selections: []ast.Selection{
			&ast.Field{
				Name: "items",
			},
			&ast.Field{
				Name: "item",
			},
			&ast.Field{
				Name: "__type",
			},
			&ast.Field{
				Name: "__schema",
			},
		},
	}

	out := ec._Query(context.Background(), sel)
	if out == graphql.Null {
		t.Error("expected non-null result")
	}

	// TODO: check the actual result
}

func Test_Query_with_error(t *testing.T) {
	ec := executionContext{
		OperationContext: &graphql.OperationContext{
			Variables: map[string]interface{}{},
		},
	}

	sel := ast.SelectionSet{
		Selections: []ast.Selection{
			&ast.Field{
				Name: "items",
			},
			&ast.Field{
				Name: "item",
			},
			&ast.Field{
				Name: "__type",
			},
			&ast.Field{
				Name: "__schema",
			},
		},
	}

	ec.Error(context.Background(), errors.New("some error"))

	out := ec._Query(context.Background(), sel)
	if out != graphql.Null {
		t.Error("expected null result")
	}
}
