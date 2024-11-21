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

func TestUnmarshalNBoolean2bool_c4af251785(t *testing.T) {
	ec := executionContext{}
	// Success case
	v := true
	got, err := ec.unmarshalNBoolean2bool(context.Background(), v)
	if err != nil {
		t.Error(err)
	}
	if got != v {
		t.Errorf("Expected %v, got %v", v, got)
	}

	// Failure case
	v = "not a boolean"
	got, err = ec.unmarshalNBoolean2bool(context.Background(), v)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if got != false {
		t.Errorf("Expected %v, got %v", false, got)
	}
}
