package generated

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"testing"
)

// Test___Type_enumValues_b274b3902b tests the ___Type_enumValues method of the execution context
func Test___Type_enumValues_b274b3902b(t *testing.T) {
	ctx := context.Background()
	ec := &executionContext{}
	obj := &introspection.Type{}
	collectedField := &graphql.CollectedField{}

	// Testing when the method successfully executes
	marshaler := ec.___Type_enumValues(ctx, collectedField, obj)
	if marshaler != graphql.Null {
		t.Log("Method successfully executed, return value not null")
	} else {
		t.Error("Method did not execute successfully, return value is null")
	}

	// Testing when obj is null and hence method returns null 
	obj = nil
	marshaler = ec.___Type_enumValues(ctx, collectedField, obj)
	if marshaler == graphql.Null {
		t.Log("Method executed as expected and returned null as obj is nil")
	} else {
		t.Error("Method did not execute as expected, obj was nil but return value is not null")
	}

	// TODO: Add more test cases to cover different scenarios and edge cases 
}
