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

func Test_Item_5a63328662(t *testing.T) {
	ec := executionContext{
		OperationContext: &graphql.OperationContext{
			Variables: map[string]interface{}{},
		},
	}

	sel := ast.SelectionSet{
		Selections: []ast.Selection{
			&ast.Field{
				Name: "id",
			},
			&ast.Field{
				Name: "title",
			},
			&ast.Field{
				Name: "owner",
			},
			&ast.Field{
				Name: "rating",
			},
		},
	}

	obj := &model.Item{
		ID:      "1",
		Title:   "Test Item",
		Owner:   "John Doe",
		Rating:  5,
	}

	out := ec._Item(context.Background(), sel, obj)
	if out == graphql.Null {
		t.Error("Expected non-null result")
	}

	fields := graphql.CollectFields(ec.OperationContext, sel, itemImplementors)
	if len(fields) != len(out.Values) {
		t.Errorf("Expected %d fields, got %d", len(fields), len(out.Values))
	}

	for i, field := range fields {
		switch field.Name {
		case "__typename":
			if out.Values[i] != graphql.MarshalString("Item") {
				t.Errorf("Expected __typename to be \"Item\", got %v", out.Values[i])
			}
		case "id":
			if out.Values[i] != graphql.MarshalString("1") {
				t.Errorf("Expected id to be \"1\", got %v", out.Values[i])
			}
		case "title":
			if out.Values[i] != graphql.MarshalString("Test Item") {
				t.Errorf("Expected title to be \"Test Item\", got %v", out.Values[i])
			}
		case "owner":
			if out.Values[i] != graphql.MarshalString("John Doe") {
				t.Errorf("Expected owner to be \"John Doe\", got %v", out.Values[i])
			}
		case "rating":
			if out.Values[i] != graphql.MarshalInt(5) {
				t.Errorf("Expected rating to be 5, got %v", out.Values[i])
			}
		}
	}
}

func Test_Item_Error(t *testing.T) {
	ec := executionContext{
		OperationContext: &graphql.OperationContext{
			Variables: map[string]interface{}{},
		},
	}

	sel := ast.SelectionSet{
		Selections: []ast.Selection{
			&ast.Field{
				Name: "id",
			},
			&ast.Field{
				Name: "title",
			},
			&ast.Field{
				Name: "owner",
			},
			&ast.Field{
				Name: "rating",
			},
		},
	}

	obj := &model.Item{
		ID:      "1",
		Title:   "Test Item",
		Owner:   "John Doe",
		Rating:  5,
	}

	ec.ResolverMiddleware = func(ctx context.Context, next graphql.Resolver) (res graphql.Marshaler) {
		return graphql.Null
	}

	out := ec._Item(context.Background(), sel, obj)
	if out != graphql.Null {
		t.Error("Expected null result")
	}
}
