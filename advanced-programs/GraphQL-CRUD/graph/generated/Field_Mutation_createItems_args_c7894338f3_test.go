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

func TestField_Mutation_createItems_args_c7894338f3(t *testing.T) {
	ec := executionContext{
		introspection.Schema{},
		map[string]graphql.Resolver{},
		bytes.NewBuffer(nil),
		&sync.Mutex{},
		&atomic.Value{},
	}

	// Success case
	rawArgs := map[string]interface{}{
		"input": map[string]interface{}{
			"name":  "Item 1",
			"price": 10.0,
		},
	}
	args, err := ec.field_Mutation_createItems_args(context.Background(), rawArgs)
	if err != nil {
		t.Error(err)
	}
	if args["input"].(model.NewItem).Name != "Item 1" {
		t.Error("Expected name to be Item 1")
	}
	if args["input"].(model.NewItem).Price != 10.0 {
		t.Error("Expected price to be 10.0")
	}

	// Error case: missing name
	rawArgs = map[string]interface{}{
		"input": map[string]interface{}{
			"price": 10.0,
		},
	}
	args, err = ec.field_Mutation_createItems_args(context.Background(), rawArgs)
	if err == nil {
		t.Error("Expected error for missing name")
	}
	if args != nil {
		t.Error("Expected args to be nil")
	}

	// Error case: invalid price
	rawArgs = map[string]interface{}{
		"input": map[string]interface{}{
			"name":  "Item 1",
			"price": "invalid",
		},
	}
	args, err = ec.field_Mutation_createItems_args(context.Background(), rawArgs)
	if err == nil {
		t.Error("Expected error for invalid price")
	}
	if args != nil {
		t.Error("Expected args to be nil")
	}
}
