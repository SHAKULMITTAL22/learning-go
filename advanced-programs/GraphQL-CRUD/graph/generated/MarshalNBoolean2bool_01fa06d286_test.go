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

func TestMarshalNBoolean2bool_01fa06d286(t *testing.T) {
	ec := executionContext{
		responseWriter: &bytes.Buffer{},
		introspection:  introspection.NewSchema(gqlparser.MustLoadSchema(&ast.Source{Name: "schema.graphql"})),
	}
	ctx := context.Background()
	sel := &ast.SelectionSet{
		Selections: []*ast.Selection{
			&ast.Field{
				Name: "id",
				Type: &ast.NamedType{
					Name: "ID",
				},
			},
		},
	}
	v := true
	res := ec.marshalNBoolean2bool(ctx, sel, v)
	if res == graphql.Null {
		t.Error("must not be null")
	}
}

func TestMarshalNBoolean2bool_01fa06d286_Error(t *testing.T) {
	ec := executionContext{
		responseWriter: &bytes.Buffer{},
		introspection:  introspection.NewSchema(gqlparser.MustLoadSchema(&ast.Source{Name: "schema.graphql"})),
	}
	ctx := context.Background()
	sel := &ast.SelectionSet{
		Selections: []*ast.Selection{
			&ast.Field{
				Name: "id",
				Type: &ast.NamedType{
					Name: "ID",
				},
			},
		},
	}
	v := false
	res := ec.marshalNBoolean2bool(ctx, sel, v)
	if res != graphql.Null {
		t.Error("must be null")
	}
}
