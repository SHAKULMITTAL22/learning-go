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

func TestUnmarshalN__DirectiveLocation2string_656151c8f4(t *testing.T) {
	ec := executionContext{
		unmarshalN__DirectiveLocation2string: func(ctx context.Context, v interface{}) (string, error) {
			return graphql.UnmarshalString(v)
		},
	}

	// Success case
	v := "__DirectiveLocation"
	got, err := ec.unmarshalN__DirectiveLocation2string(context.Background(), v)
	if err != nil {
		t.Error(err)
	}
	if got != "__DirectiveLocation" {
		t.Errorf("Expected __DirectiveLocation, got %s", got)
	}

	// Failure case
	v = "invalid"
	got, err = ec.unmarshalN__DirectiveLocation2string(context.Background(), v)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if got != "" {
		t.Errorf("Expected empty string, got %s", got)
	}
}
