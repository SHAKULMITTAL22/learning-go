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

func TestNewExecutableSchema_1f81986322(t *testing.T) {
	cfg := Config{
		Resolvers:  &mockResolver{},
		Directives: nil,
		Complexity: nil,
	}

	// Success case
	schema := NewExecutableSchema(cfg)
	if schema == nil {
		t.Error("NewExecutableSchema returned nil")
	}

	// Failure case
	cfg.Resolvers = nil
	schema = NewExecutableSchema(cfg)
	if schema != nil {
		t.Error("NewExecutableSchema did not return nil when resolvers were nil")
	}
}
