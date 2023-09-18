package generated

import (
	"errors"
	"testing"

	"github.com/99designs/gqlgen/graphql/introspection"
)

type ExecutionContext struct {
	DisableIntrospection bool
}

func (ec *ExecutionContext) IntrospectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(parsedSchema, parsedSchema.Types[name]), nil
}

func TestIntrospectType_d788815c23(t *testing.T) {
	ec := &ExecutionContext{}

	// Test case where introspection is disabled
	ec.DisableIntrospection = true
	_, err := ec.IntrospectType("Test")
	if err == nil {
		t.Error("Expected error, got nil")
	} else {
		t.Logf("Test Passed, Expected error and got error: %s", err.Error())
	}

	// Test case where introspection is enabled but type does not exist
	ec.DisableIntrospection = false
	_, err = ec.IntrospectType("NonExistingType")
	if err == nil {
		t.Error("Expected error, got nil")
	} else {
		t.Logf("Test Passed, Expected error and got error: %s", err.Error())
	}

	// Test case where introspection is enabled and type exists
	// TODO: Add a type to parsedSchema before testing this case
	// _, err = ec.IntrospectType("ExistingType")
	// if err != nil {
	// 	t.Errorf("Expected no error, got: %s", err.Error())
	// } else {
	// 	t.Log("Test Passed, Expected no error and got no error")
	// }
}
