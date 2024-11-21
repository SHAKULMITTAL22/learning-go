package generated

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
	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/graph/model"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// Test case method for func (ec *executionContext) ___Directive_locations.
func Test___Directive_locations_2d6c7100e1(t *testing.T) {
    // Prepare object ec
    ec := &executionContext{}

    // Prepare object of directive with valid locations
    directiveObj := &introspection.Directive{
        Locations: []string{"Location1", "Location2"},
    }

    // Prepare object of directive without locations
    directiveObjNoLocation := &introspection.Directive{}

    // Prepare dummy graphql.CollectedField object
    field := graphql.CollectedField{}

    // context
    ctx := context.Background()

    // Test case: Positive scenario, expected to return locations
    got := ec.___Directive_locations(ctx, field, directiveObj)
    expected := directiveObj.Locations
    if !bytes.Equal(got, expected) {
        t.Errorf("___Directive_locations(ctx, field, directiveObj) = %v; want %v", got, expected)
    } else {
        t.Log("Test___Directive_locations_2d6c7100e1 passed for test case: Positive scenario, expected to return locations")
    }

    // Test case: Negative scenario, no locations in input object
    got = ec.___Directive_locations(ctx, field, directiveObjNoLocation)
    expected = graphql.Null
    if !bytes.Equal(got, expected) {
        t.Errorf("___Directive_locations(ctx, field, directiveObjNoLocation) = %v; want %v", got, expected)
    } else {
        t.Log("Test___Directive_locations_2d6c7100e1 passed for test case: Negative scenario, no locations in input object")
    }
}
