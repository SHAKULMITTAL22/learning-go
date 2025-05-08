
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

func TestIntrospectType(t *testing.T) {
	t.Run("Test with introspection enabled", func(t *testing.T) {
		ec := &ExecutionContext{
			DisableIntrospection: false,
		}
		result, err := ec.IntrospectType("Query")
		if err != nil {
			t.Error("Expected no error, got ", err)
		}
		if result == nil {
			t.Error("Expected result, got null")
		}
		t.Log("Success: Test with introspection enabled")
	})

	t.Run("Test with introspection disabled", func(t *testing.T) {
		ec := &ExecutionContext{
			DisableIntrospection: true,
		}
		result, err := ec.IntrospectType("Query")
		if err == nil {
			t.Error("Expected error, got none")
		}
		if result != nil {
			t.Error("Expected no result, got ", result)
		}
		t.Log("Success: Test with introspection disabled")
	})

	t.Run("Test with non-existing type", func(t *testing.T) {
		ec := &ExecutionContext{
			DisableIntrospection: false,
		}
		result, err := ec.IntrospectType("NonExistent")
		if err != nil {
			t.Error("Expected no error, got ", err)
		}
		if result != nil {
			t.Error("Expected no result, got ", result)
		}
		t.Log("Success: Test with non-existing type")
	})
}

