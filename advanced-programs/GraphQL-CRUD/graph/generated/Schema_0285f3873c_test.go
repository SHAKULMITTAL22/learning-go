package generated

import (
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/graph/model"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

func TestSchema_0285f3873c(t *testing.T) {
	parsedSchema := &ast.Schema{}

	e := &executableSchema{
		Schema: func() *ast.Schema {
			return parsedSchema
		},
	}

	// Test case 1: Valid schema
	schema, err := e.Schema()
	if err != nil {
		t.Error(err)
	}
	if schema == nil {
		t.Error("Schema should not be nil")
	}

	// Test case 2: Invalid schema
	parsedSchema = nil
	schema, err = e.Schema()
	if err == nil {
		t.Error("Error should not be nil")
	}
	if schema != nil {
		t.Error("Schema should be nil")
	}
}
