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

func TestField_Mutation_delete_args_af684aa666(t *testing.T) {
	ec := executionContext{
		unmarshalNInt2int: func(ctx context.Context, v interface{}) (int, error) {
			return 1, nil
		},
	}

	// Success case
	rawArgs := map[string]interface{}{
		"id": 1,
	}
	args, err := ec.field_Mutation_delete_args(context.Background(), rawArgs)
	if err != nil {
		t.Error(err)
	}
	if args["id"] != 1 {
		t.Error("Expected args[id] to be 1, but got", args["id"])
	}

	// Error case
	rawArgs = map[string]interface{}{
		"id": "not an int",
	}
	args, err = ec.field_Mutation_delete_args(context.Background(), rawArgs)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if args != nil {
		t.Error("Expected args to be nil, but got", args)
	}
}
