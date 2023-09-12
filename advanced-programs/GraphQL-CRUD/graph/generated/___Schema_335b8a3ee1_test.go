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

func Test___Schema_335b8a3ee1(t *testing.T) {
	ec := executionContext{
		OperationContext: &graphql.OperationContext{},
	}
	sel := ast.SelectionSet{
		Fields: []*ast.Field{
			{
				Name: "__typename",
			},
			{
				Name: "types",
			},
			{
				Name: "queryType",
			},
			{
				Name: "mutationType",
			},
			{
				Name: "subscriptionType",
			},
			{
				Name: "directives",
			},
		},
	}
	obj := &introspection.Schema{}
	out := ec.___Schema(context.Background(), sel, obj)
	if out == graphql.Null {
		t.Error("expected non-null result")
	}
	fields := graphql.CollectFields(ec.OperationContext, sel, __SchemaImplementors)
	if len(fields) != len(out.Values) {
		t.Errorf("expected %d fields, got %d", len(fields), len(out.Values))
	}
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			if out.Values[i] != graphql.MarshalString("__Schema") {
				t.Errorf("expected __typename to be __Schema, got %s", out.Values[i])
			}
		case "types":
			if out.Values[i] == graphql.Null {
				t.Error("expected non-null result")
			}
		case "queryType":
			if out.Values[i] == graphql.Null {
				t.Error("expected non-null result")
			}
		case "mutationType":
			if out.Values[i] == graphql.Null {
				t.Error("expected non-null result")
			}
		case "subscriptionType":
			if out.Values[i] == graphql.Null {
				t.Error("expected non-null result")
			}
		case "directives":
			if out.Values[i] == graphql.Null {
				t.Error("expected non-null result")
			}
		default:
			t.Errorf("unknown field %s", field.Name)
		}
	}
}

func Test___Schema_335b8a3ee1_WithErrors(t *testing.T) {
	ec := executionContext{
		OperationContext: &graphql.OperationContext{},
	}
	sel := ast.SelectionSet{
		Fields: []*ast.Field{
			{
				Name: "__typename",
			},
			{
				Name: "types",
			},
			{
				Name: "queryType",
			},
			{
				Name: "mutationType",
			},
			{
				Name: "subscriptionType",
			},
			{
				Name: "directives",
			},
		},
	}
	obj := &introspection.Schema{}
	out := ec.___Schema(context.Background(), sel, obj)
	if out == graphql.Null {
		t.Error("expected non-null result")
	}
	fields := graphql.CollectFields(ec.OperationContext, sel, __SchemaImplementors)
	if len(fields) != len(out.Values) {
		t.Errorf("expected %d fields, got %d", len(fields), len(out.Values))
	}
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			if out.Values[i] != graphql.MarshalString("__Schema") {
				t.Errorf("expected __typename to be __Schema, got %s", out.Values[i])
			}
		case "types":
			if out.Values[i] == graphql.Null {
				t.Error("expected non-null result")
			}
		case "queryType":
			if out.Values[i] == graphql.Null {
				t.Error("expected non-null result")
			}
		case "mutationType":
			if out.Values[i] == graphql.Null {
				t.Error("expected non-null result")
			}
		case "subscriptionType":
			if out.Values[i] == graphql.Null {
				t.Error("expected non-null result")
			}
		case "directives":
			if out.Values[i] == graphql.Null {
				t.Error("expected non-null result")
			}
		default:
			t.Errorf("unknown field %s", field.Name)
		}
	}
}
