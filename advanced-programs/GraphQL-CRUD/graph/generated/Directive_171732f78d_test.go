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

func TestDirective_171732f78d(t *testing.T) {
	ec := executionContext{
		directives: make(map[string]*introspection.Directive),
	}

	// Success case
	v := introspection.Directive{
		Name: "testDirective",
		Args: []introspection.InputValue{
			{
				Name: "arg1",
				Type: "String",
			},
			{
				Name: "arg2",
				Type: "Int",
			},
		},
	}
	marshaler := ec.marshalN__Directive2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐDirective(context.Background(), nil, v)
	if marshaler == nil {
		t.Error("marshaler should not be nil")
	}

	// Failure case: directive not found
	v = introspection.Directive{
		Name: "nonexistentDirective",
	}
	marshaler = ec.marshalN__Directive2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐDirective(context.Background(), nil, v)
	if marshaler != nil {
		t.Error("marshaler should be nil")
	}

	// Failure case: invalid argument type
	v = introspection.Directive{
		Name: "testDirective",
		Args: []introspection.InputValue{
			{
				Name: "arg1",
				Type: "String",
			},
			{
				Name: "arg2",
				Type: "Float", // Invalid type
			},
		},
	}
	marshaler = ec.marshalN__Directive2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐDirective(context.Background(), nil, v)
	if marshaler != nil {
		t.Error("marshaler should be nil")
	}
}
