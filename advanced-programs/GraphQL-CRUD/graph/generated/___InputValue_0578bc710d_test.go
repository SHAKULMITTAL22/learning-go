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

func Test___InputValue_0578bc710d(t *testing.T) {
	ec := executionContext{
		OperationContext: &graphql.OperationContext{
			Variables: map[string]interface{}{},
		},
	}

	sel := ast.SelectionSet{
		Fields: []*ast.Field{
			{
				Name: "__typename",
			},
			{
				Name: "name",
			},
			{
				Name: "description",
			},
			{
				Name: "type",
			},
			{
				Name: "defaultValue",
			},
		},
	}

	obj := &introspection.InputValue{
		Name:        "name",
		Description: "description",
		Type:        "type",
		DefaultValue: "defaultValue",
	}

	out := ec.___InputValue(context.Background(), sel, obj)

	if out == graphql.Null {
		t.Error("expected non-null value")
	}

	if out.String() != `{"__typename":"__InputValue","name":"name","description":"description","type":"type","defaultValue":"defaultValue"}` {
		t.Errorf("expected output to be `%s`, got `%s`", `{"__typename":"__InputValue","name":"name","description":"description","type":"type","defaultValue":"defaultValue"}`, out.String())
	}
}

func Test___InputValue_with_invalid_field(t *testing.T) {
	ec := executionContext{
		OperationContext: &graphql.OperationContext{
			Variables: map[string]interface{}{},
		},
	}

	sel := ast.SelectionSet{
		Fields: []*ast.Field{
			{
				Name: "__typename",
			},
			{
				Name: "invalid_field",
			},
		},
	}

	obj := &introspection.InputValue{
		Name:        "name",
		Description: "description",
		Type:        "type",
		DefaultValue: "defaultValue",
	}

	out := ec.___InputValue(context.Background(), sel, obj)

	if out != graphql.Null {
		t.Error("expected null value")
	}
}
