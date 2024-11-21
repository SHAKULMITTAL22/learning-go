package generated

import (
	"bytes"
	"context"
	"errors"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/vektah/gqlparser/v2/ast"
)

// Testing ___Type method
func Test___Type_729d6aeb56(t *testing.T) {
	ctx := context.Background()
	fields := []*ast.Field{
		{Name: "kind"},
		{Name: "name"},
		{Name: "invalid_field"},
	}

	// Test case: when the field name is invalid
	ec := &executionContext{}
	obj := &introspection.Type{}
	
	// Expect a panic when the field name is invalid
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Method did not panic with invalid field name")
		}
	}()

	ec.___Type(ctx, fields, obj)

	// Test case: when the field's value is null
	ec = &executionContext{
		___Type_kind: func(context.Context, *ast.Field, *introspection.Type) graphql.Marshaler {
			return graphql.Null
		},
	}
	obj = &introspection.Type{}

	result := ec.___Type(ctx, fields, obj)
	if _, ok := result.(graphql.Null); !ok {
		t.Errorf(
			"Expected result type graphql.Null, got %T. args: [ctx: %v, fields: %v, obj: %v]",
			result, ctx, fields, obj,
		)
	}
}
