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

func Test___Directive_2103688df7(t *testing.T) {
	ec := executionContext{
		OperationContext: &graphql.OperationContext{
			Variables: map[string]interface{}{},
		},
	}

	sel := ast.SelectionSet{
		Selections: []ast.Selection{
			&ast.Field{
				Name: "__typename",
			},
			&ast.Field{
				Name: "name",
			},
			&ast.Field{
				Name: "description",
			},
			&ast.Field{
				Name: "locations",
			},
			&ast.Field{
				Name: "args",
			},
		},
	}

	obj := &introspection.Directive{
		Name:        "directiveName",
		Description: "directive description",
		Locations:   []string{"QUERY", "MUTATION"},
		Args: []*introspection.InputValue{
			{
				Name:        "argName",
				Description: "arg description",
				Type:        "String",
			},
		},
	}

	out := ec.___Directive(context.Background(), sel, obj)
	if out == graphql.Null {
		t.Error("expected non-null result")
		return
	}

	fields := out.(*graphql.FieldSet).Fields
	if len(fields) != 6 {
		t.Errorf("expected 6 fields, got %d", len(fields))
		return
	}

	// Check __typename field
	if fields[0].Name != "__typename" {
		t.Errorf("expected __typename field to be first, got %s", fields[0].Name)
		return
	}
	if fields[0].Value.(string) != "__Directive" {
		t.Errorf("expected __typename field to be __Directive, got %s", fields[0].Value.(string))
		return
	}

	// Check name field
	if fields[1].Name != "name" {
		t.Errorf("expected name field to be second, got %s", fields[1].Name)
		return
	}
	if fields[1].Value.(string) != "directiveName" {
		t.Errorf("expected name field to be directiveName, got %s", fields[1].Value.(string))
		return
	}

	// Check description field
	if fields[2].Name != "description" {
		t.Errorf("expected description field to be third, got %s", fields[2].Name)
		return
	}
	if fields[2].Value.(string) != "directive description" {
		t.Errorf("expected description field to be directive description, got %s", fields[2].Value.(string))
		return
	}

	// Check locations field
	if fields[3].Name != "locations" {
		t.Errorf("expected locations field to be fourth, got %s", fields[3].Name)
		return
	}
	if len(fields[3].Value.([]string)) != 2 {
		t.Errorf("expected locations field to have 2 elements, got %d", len(fields[3].Value.([]string)))
		return
	}
	if fields[3].Value.([]string)[0] != "QUERY" {
		t.Errorf("expected locations field to contain QUERY, got %s", fields[3].Value.([]string)[0])
		return
	}
	if fields[3].Value.([]string)[1] != "MUTATION" {
		t.Errorf("expected locations field to contain MUTATION, got %s", fields[3].Value.([]string)[1])
		return
	}

	// Check args field
	if fields[4].Name != "args" {
		t.Errorf("expected args field to be fifth, got %s", fields[4].Name)
		return
	}
	if len(fields[4].Value.([]*introspection.InputValue)) != 1 {
		t.Errorf("expected args field to have 1 element, got %d", len(fields[4].Value.([]*introspection.InputValue)))
		return
	}
	if fields[4].Value.([]*introspection.InputValue)[0].Name != "argName" {
		t.Errorf("expected args field to contain argName, got %s", fields[4].Value.([]*introspection.InputValue)[0].Name)
		return
	}
	if fields[4].Value.([]*introspection.InputValue)[0].Description != "arg description" {
		t.Errorf("expected args field to contain arg description, got %s", fields[4].Value.([]*introspection.InputValue)[0].Description)
		return
	}
	if fields[4].Value.([]*introspection.InputValue)[0].Type != "String" {
		t.Errorf("expected args field to contain String, got %s", fields[4].Value.([]*introspection.InputValue)[0].Type)
		return
	}
}

func Test___Directive_Error(t *testing.T) {
	ec := executionContext{
		OperationContext: &graphql.OperationContext{
			Variables: map[string]interface{}{},
		},
	}

	sel := ast.SelectionSet{
		Selections: []ast.Selection{
			&ast.Field{
				Name: "__typename",
			},
			&ast.Field{
				Name: "name",
			},
			&ast.Field{
				Name: "description",
			},
			&ast.Field{
				Name: "locations",
			},
			&ast.Field{
				Name: "args",
			},
		},
	}

	obj := &introspection.Directive{}

	out := ec.___Directive(context.Background(), sel, obj)
	if out != graphql.Null {
		t.Error("expected null result")
		return
	}
}
