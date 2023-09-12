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

func TestUnmarshalNString2string_755d3e6071(t *testing.T) {
	ec := executionContext{unmarshalNString2string: func(ctx context.Context, v interface{}) (string, error) {
		return graphql.UnmarshalString(v)
	}}

	// Success case
	v := "Hello, world!"
	got, err := ec.unmarshalNString2string(context.Background(), v)
	if err != nil {
		t.Error(err)
	}
	if got != v {
		t.Errorf("Expected %q, got %q", v, got)
	}

	// Failure case
	v = 123
	got, err = ec.unmarshalNString2string(context.Background(), v)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if got != "" {
		t.Errorf("Expected empty string, got %q", got)
	}
}
