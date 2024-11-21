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

func TestMarshalNString2string_8f8acf14f2(t *testing.T) {
	ec := executionContext{
		introspection.Schema{},
		map[string]graphql.Marshaler{},
		map[string]graphql.Type{},
		map[string]graphql.Type{},
		bytes.NewBuffer(nil),
		&sync.Mutex{},
		&atomic.Value{},
		context.Background(),
	}
	sel := &ast.SelectionSet{
		Selections: []*ast.Selection{},
	}
	v := "string"
	res := ec.marshalNString2string(context.Background(), sel, v)
	if res == graphql.Null {
		t.Error("must not be null")
	}
}

func TestMarshalNString2string_error(t *testing.T) {
	ec := executionContext{
		introspection.Schema{},
		map[string]graphql.Marshaler{},
		map[string]graphql.Type{},
		map[string]graphql.Type{},
		bytes.NewBuffer(nil),
		&sync.Mutex{},
		&atomic.Value{},
		context.Background(),
	}
	sel := &ast.SelectionSet{
		Selections: []*ast.Selection{},
	}
	v := errors.New("error")
	res := ec.marshalNString2string(context.Background(), sel, v)
	if res != graphql.Null {
		t.Error("must be null")
	}
}
