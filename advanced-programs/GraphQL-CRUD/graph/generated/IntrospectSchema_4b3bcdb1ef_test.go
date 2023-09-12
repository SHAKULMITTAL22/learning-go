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

func TestIntrospectSchema_4b3bcdb1ef(t *testing.T) {
	ec := &executionContext{
		DisableIntrospection: false,
	}

	// Success case
	schema, err := ec.introspectSchema()
	if err != nil {
		t.Error(err)
	}
	if schema == nil {
		t.Error("schema is nil")
	}

	// Failure case
	ec.DisableIntrospection = true
	schema, err = ec.introspectSchema()
	if err == nil {
		t.Error("expected error, got nil")
	}
	if schema != nil {
		t.Error("expected nil, got schema")
	}
}
