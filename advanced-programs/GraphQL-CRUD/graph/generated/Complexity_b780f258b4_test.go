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

func TestComplexity_b780f258b4(t *testing.T) {
	ec := executionContext{nil, &executableSchema{}}
	// Success case
	complexity, ok := ec.Complexity("Item", "id", 0, nil)
	if !ok {
		t.Error("Complexity should be true")
	}
	if complexity != 0 {
		t.Error("Complexity should be 0")
	}
	// Failure case
	complexity, ok = ec.Complexity("Item", "invalidField", 0, nil)
	if ok {
		t.Error("Complexity should be false")
	}
	if complexity != 0 {
		t.Error("Complexity should be 0")
	}
}
