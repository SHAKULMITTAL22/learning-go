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

func TestExec_084531e294(t *testing.T) {
	// Create a context with an operation context.
	ctx := context.Background()
	rc := graphql.GetOperationContext(ctx)
	rc.Operation = &ast.OperationDefinition{
		Operation: ast.Query,
		SelectionSet: &ast.SelectionSet{
			Selections: []ast.Selection{
				&ast.Field{
					Name: &ast.Name{
						Value: "query",
					},
				},
			},
		},
	}

	// Create an executable schema.
	e := &executableSchema{
		rc: rc,
	}

	// Execute the query.
	resp := e.Exec(ctx)

	// Check the response.
	if resp.Data == nil {
		t.Error("expected data to be non-nil")
	}

	// Check the response data.
	var data map[string]interface{}
	if err := json.Unmarshal(resp.Data, &data); err != nil {
		t.Error("expected data to be valid JSON")
	}

	// Check the response data for the expected fields.
	if _, ok := data["query"]; !ok {
		t.Error("expected data to contain the 'query' field")
	}
}

func TestExec_Error(t *testing.T) {
	// Create a context with an operation context.
	ctx := context.Background()
	rc := graphql.GetOperationContext(ctx)
	rc.Operation = &ast.OperationDefinition{
		Operation: ast.Query,
		SelectionSet: &ast.SelectionSet{
			Selections: []ast.Selection{
				&ast.Field{
					Name: &ast.Name{
						Value: "query",
					},
				},
			},
		},
	}

	// Create an executable schema.
	e := &executableSchema{
		rc: rc,
	}

	// Execute the query.
	resp := e.Exec(ctx)

	// Check the response.
	if resp.Data != nil {
		t.Error("expected data to be nil")
	}

	// Check the response errors.
	if len(resp.Errors) == 0 {
		t.Error("expected errors to be non-empty")
	}

	// Check the response errors for the expected error.
	if resp.Errors[0].Message != "unsupported GraphQL operation" {
		t.Error("expected error to be 'unsupported GraphQL operation'")
	}
}
