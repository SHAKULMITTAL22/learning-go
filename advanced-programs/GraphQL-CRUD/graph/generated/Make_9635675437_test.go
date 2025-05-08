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

func TestMake_9635675437(t *testing.T) {
	ec := executionContext{
		introspection: introspection.Schema{},
	}
	sel := ast.SelectionSet{
		Fields: []*ast.Field{
			{
				Name: "name",
			},
		},
	}
	v := []introspection.InputValue{
		{
			Name: "name",
			Type: "String",
		},
	}
	ret := ec.marshalN__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValueᚄ(context.Background(), sel, v)
	if ret == nil {
		t.Error("expected non-nil return value")
	}
	if len(ret) != len(v) {
		t.Errorf("expected return value to have length %d, got %d", len(v), len(ret))
	}
	for i, r := range ret {
		if r == nil {
			t.Errorf("expected non-nil return value at index %d", i)
		}
		if r.String() != v[i].Name {
			t.Errorf("expected return value at index %d to be %s, got %s", i, v[i].Name, r.String())
		}
	}
}

func TestMake_9635675437_Error(t *testing.T) {
	ec := executionContext{
		introspection: introspection.Schema{},
	}
	sel := ast.SelectionSet{
		Fields: []*ast.Field{
			{
				Name: "name",
			},
		},
	}
	v := []introspection.InputValue{
		{
			Name: "name",
			Type: "String",
		},
		{
			Name: "age",
			Type: "Int",
		},
	}
	ret := ec.marshalN__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValueᚄ(context.Background(), sel, v)
	if ret != nil {
		t.Error("expected nil return value")
	}
}
