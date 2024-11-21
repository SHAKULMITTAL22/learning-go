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

func TestMake_56fead5ac9(t *testing.T) {
	ec := executionContext{
		directives: make(map[string]*ast.Directive),
		variables: make(map[string]interface{}),
	}

	// Success case
	v := []interface{}{"foo", "bar"}
	res, err := ec.unmarshalN__DirectiveLocation2ᚕstringᚄ(context.Background(), v)
	if err != nil {
		t.Error(err)
	}
	if len(res) != 2 {
		t.Error("Expected 2 results, got", len(res))
	}
	if res[0] != "foo" || res[1] != "bar" {
		t.Error("Expected foo and bar, got", res)
	}

	// Error case
	v = []interface{}{"foo", 123}
	_, err = ec.unmarshalN__DirectiveLocation2ᚕstringᚄ(context.Background(), v)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if !errors.Is(err, gqlparser.ErrInvalidType) {
		t.Error("Expected ErrInvalidType, got", err)
	}
}
