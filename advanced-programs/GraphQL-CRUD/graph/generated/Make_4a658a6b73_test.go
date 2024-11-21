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

func TestMake_4a658a6b73(t *testing.T) {
	ec := executionContext{
		directives: make(map[string]*ast.Directive),
		variables:  make(map[string]interface{}),
	}
	ctx := context.Background()
	sel := &ast.SelectionSet{
		Selections: []*ast.Selection{
			&ast.Field{
				Name: "directiveLocation",
			},
		},
	}
	v := []string{"QUERY"}
	graphql.AddError(ctx, errors.New("error"))
	ec.marshalN__DirectiveLocation2ᚕstringᚄ(ctx, sel, v)
	if len(ec.Errors) == 0 {
		t.Error("expected error")
	}
}

func TestMake_4a658a6b73_2(t *testing.T) {
	ec := executionContext{
		directives: make(map[string]*ast.Directive),
		variables:  make(map[string]interface{}),
	}
	ctx := context.Background()
	sel := &ast.SelectionSet{
		Selections: []*ast.Selection{
			&ast.Field{
				Name: "directiveLocation",
			},
		},
	}
	v := []string{"QUERY"}
	ec.marshalN__DirectiveLocation2ᚕstringᚄ(ctx, sel, v)
	if len(ec.Errors) != 0 {
		t.Error("unexpected error")
	}
}
