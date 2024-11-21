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

func TestField___Type_enumValues_args_6a87c7ce87(t *testing.T) {
	ec := executionContext{
		introspection.Schema{},
		map[string]graphql.Type{},
		map[string]graphql.Type{},
		&sync.Mutex{},
		bytes.NewBuffer(nil),
		&atomic.Value{},
		nil,
		context.Background(),
		map[string]interface{}{},
		map[string]interface{}{},
	}

	// Success case
	rawArgs := map[string]interface{}{
		"includeDeprecated": true,
	}
	args, err := ec.field___Type_enumValues_args(context.Background(), rawArgs)
	if err != nil {
		t.Error(err)
	}
	if args["includeDeprecated"] != true {
		t.Error("Expected args[\"includeDeprecated\"] to be true")
	}

	// Error case
	rawArgs = map[string]interface{}{
		"includeDeprecated": "not a boolean",
	}
	args, err = ec.field___Type_enumValues_args(context.Background(), rawArgs)
	if err == nil {
		t.Error("Expected error when passing non-boolean value to \"includeDeprecated\"")
	}
	if args != nil {
		t.Error("Expected args to be nil when error is returned")
	}
}
