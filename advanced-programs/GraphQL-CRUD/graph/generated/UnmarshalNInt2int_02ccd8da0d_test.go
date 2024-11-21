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

func TestUnmarshalNInt2int_02ccd8da0d(t *testing.T) {
	ec := executionContext{}
	// Success case
	v := 123
	got, err := ec.unmarshalNInt2int(context.Background(), v)
	if err != nil {
		t.Error(err)
	}
	if got != 123 {
		t.Errorf("Expected 123, got %d", got)
	}

	// Error case
	v = "abc"
	got, err = ec.unmarshalNInt2int(context.Background(), v)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if got != 0 {
		t.Errorf("Expected 0, got %d", got)
	}
}
