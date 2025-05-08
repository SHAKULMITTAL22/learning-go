// This is Go test code.

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
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
	"testing"
)

// TestMake_1607afee73 unit tests the marshalO__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐTypeᚄ method
func TestMake_1607afee73(t *testing.T) {
	// Generate a selection set to use in the function
	selSet := ast.SelectionSet{}
	ctx := context.Background()

	ec := &executionContext{}

	// Case when 'v' is nil.
	// The function should return graphql.Null
	v := []introspection.Type(nil)
	result := ec.marshalO__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐTypeᚄ(ctx, selSet, v)
	if result != graphql.Null {
		t.Error("Failed: Expected graphql.Null but got ", result)
	} else {
		t.Log("Success: Expected graphql.Null and got ", result)
	}

	// Case when 'v' is nil
	// Here, it is expected to return a not nil value
	v = []introspection.Type{
		introspection.Type{
			Name: "introType",
		},
	}
	result = ec.marshalO__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐTypeᚄ(ctx, selSet, v)
	if result == graphql.Null || result == nil {
		t.Error("Failed: Expected not nil value but got ", result)
	} else {
		t.Log("Success: Expected not nil value and got ", result)
	}
}
