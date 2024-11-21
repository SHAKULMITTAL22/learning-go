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

func TestIntrospectType_d788815c23(t *testing.T) {
	ec := executionContext{
		DisableIntrospection: false,
		parsedSchema:        parsedSchema,
	}

	// Success case
	typeArgs := make(map[string]interface{})
	typeArgs["name"] = "Query"
	result, err := ec.introspectType("Query", typeArgs)
	if err != nil {
		t.Error(err)
	}
	if result == nil {
		t.Error("introspectType returned nil")
	}
	if result.Name != "Query" {
		t.Errorf("introspectType returned incorrect type name: %s", result.Name)
	}

	// Error case
	typeArgs = make(map[string]interface{})
	typeArgs["name"] = "InvalidType"
	result, err = ec.introspectType("InvalidType", typeArgs)
	if err == nil {
		t.Error("introspectType did not return an error for an invalid type")
	}
	if result != nil {
		t.Errorf("introspectType returned a non-nil result for an invalid type: %s", result.Name)
	}
}
