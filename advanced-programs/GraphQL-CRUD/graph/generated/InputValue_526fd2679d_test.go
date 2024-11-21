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

func TestInputValue_526fd2679d(t *testing.T) {
	ec := executionContext{
		introspection.InputValue{
			Name: "name",
			Type: &ast.NamedType{
				Name: "String",
			},
		},
	}

	// Success case
	buf := bytes.NewBuffer(nil)
	err := ec.marshalN__InputValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(context.Background(), nil, &introspection.InputValue{
		Name: "name",
		Type: &ast.NamedType{
			Name: "String",
		},
	})
	if err != nil {
		t.Error(err)
	}
	if buf.String() != `"name"` {
		t.Errorf("Expected `name`, got %s", buf.String())
	}

	// Error case
	buf = bytes.NewBuffer(nil)
	err = ec.marshalN__InputValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(context.Background(), nil, &introspection.InputValue{
		Name: "name",
		Type: &ast.NamedType{
			Name: "Int",
		},
	})
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if buf.String() != "" {
		t.Errorf("Expected empty string, got %s", buf.String())
	}
}
