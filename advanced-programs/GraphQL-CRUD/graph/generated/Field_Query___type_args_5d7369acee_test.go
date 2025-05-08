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

func TestField_Query___type_args_5d7369acee(t *testing.T) {
	ec := executionContext{
		unmarshalNString2string: func(ctx context.Context, v interface{}) (string, error) {
			return strconv.Unquote(v.(string))
		},
	}

	// Success case
	rawArgs := map[string]interface{}{
		"name": "MyType",
	}
	args, err := ec.field_Query___type_args(context.Background(), rawArgs)
	if err != nil {
		t.Error(err)
	}
	if args["name"] != "MyType" {
		t.Error("Expected args['name'] to be 'MyType', but got", args["name"])
	}

	// Error case
	rawArgs = map[string]interface{}{
		"name": 123,
	}
	args, err = ec.field_Query___type_args(context.Background(), rawArgs)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if !errors.Is(err, gqlparser.ErrInvalidType) {
		t.Error("Expected error to be gqlparser.ErrInvalidType, but got", err)
	}
}
