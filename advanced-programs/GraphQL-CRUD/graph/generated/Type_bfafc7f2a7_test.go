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

func TestType_bfafc7f2a7(t *testing.T) {
	ec := executionContext{
		introspection.Type{},
	}
	sel := ast.SelectionSet{}
	v := introspection.Type{}
	graphql.Marshaler{}

	// Success case
	ec.marshalN__Type2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(context.Background(), sel, v)

	// Failure case
	ec.marshalN__Type2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(context.Background(), sel, nil)
	if !t.Failed() {
		t.Error("Expected error when marshaling nil value")
	}
}
