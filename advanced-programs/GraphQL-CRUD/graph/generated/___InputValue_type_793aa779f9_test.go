package generated

import (
	"context"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/stretchr/testify/assert"
)

// Test___InputValue_type_793aa779f9 is a unit test for ___InputValue_type
func Test___InputValue_type_793aa779f9(t *testing.T) {
	// Creating an execution context 
	executionContext := graphql.ExecutionContext{}

	// Setting up a field 
	field := graphql.CollectedField{}

	inputValueType := "String"

	inputValue := introspection.InputValue{
		Type: &introspection.Type{
			Name: &inputValueType,
		},
	}

	ctx := context.Background()

	// Successful test case where the return value should not be null
	marshalObj := executionContext.___InputValue_type(ctx, &field, &inputValue)
	marshalRes, _ := marshalObj.MarshalJSON()
	assert.NotEqual(t, string(marshalRes), "null")

	// Failure test where a panic will occur if the input value is null
	inputValueNull := introspection.InputValue{
		Type: nil,
	}
	marshalObjFail := executionContext.___InputValue_type(ctx, &field, &inputValueNull)
	marshalResFail, _ := marshalObjFail.MarshalJSON()
	assert.Equal(t, string(marshalResFail), "null")
}
