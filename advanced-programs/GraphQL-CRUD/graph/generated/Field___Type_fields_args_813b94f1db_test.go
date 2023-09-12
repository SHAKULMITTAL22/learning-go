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

func TestField___Type_fields_args_813b94f1db(t *testing.T) {
	ec := executionContext{
		introspection.Schema{},
		map[string]graphql.ExecutableSchema{},
		bytes.NewBufferString(""),
		&sync.Mutex{},
		&atomic.Value{},
		context.Background(),
	}
	rawArgs := map[string]interface{}{
		"includeDeprecated": true,
	}
	args, err := ec.field___Type_fields_args(context.Background(), rawArgs)
	if err != nil {
		t.Error(err)
	}
	if args["includeDeprecated"] != true {
		t.Error("expected args[includeDeprecated] to be true")
	}
}

func TestField___Type_fields_args_with_invalid_arg(t *testing.T) {
	ec := executionContext{
		introspection.Schema{},
		map[string]graphql.ExecutableSchema{},
		bytes.NewBufferString(""),
		&sync.Mutex{},
		&atomic.Value{},
		context.Background(),
	}
	rawArgs := map[string]interface{}{
		"includeDeprecated": "invalid",
	}
	args, err := ec.field___Type_fields_args(context.Background(), rawArgs)
	if err == nil {
		t.Error("expected error")
	}
	if args != nil {
		t.Error("expected args to be nil")
	}
}
