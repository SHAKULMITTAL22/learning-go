package generated_test

import (
	"bytes"
	"context"
	"errors"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	generated "github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/graph/generated"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

func TestMake_2f4ae983df(t *testing.T) {
	ec := &generated.ExecutionContext{}
	sel := ast.SelectionSet{}
	ctx := context.Background()

	// Test with single introspection.Type
	v := []introspection.Type{{Name: "OnlyOneType"}}
	ret := ec.MarshalN__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐTypeᚄ(ctx, sel, v)
	if len(ret) != 1 {
		t.Error("Expected array of length 1, got ", len(ret))
	}
	t.Log("Passed test with single introspection.Type")

	// Test with multiple introspection.Types
	v = []introspection.Type{{Name: "Type1"}, {Name: "Type2"}}
	ret = ec.MarshalN__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐTypeᚄ(ctx, sel, v)
	if len(ret) != 2 {
		t.Error("Expected array of length 2, got ", len(ret))
	}
	t.Log("Passed test with multiple introspection.Types")

	// Test with nil introspection.Type in array
	v = []introspection.Type{nil}
	defer func() {
		if r := recover(); r != nil {
			t.Log("Recovered from panic as expected")
		}
	}()
	ret = ec.MarshalN__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐTypeᚄ(ctx, sel, v)
	if ret != nil {
		t.Error("Expected nil array due to panic, got ", ret)
	}
	t.Log("Passed test with nil introspection.Type in array")
}
