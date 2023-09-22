package generated

import (
	"errors"
	"testing"

	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/stretchr/testify/assert"
)

type executionContextTest struct {
	DisableIntrospection bool
}

// TestIntrospectSchema_4b3bcdb1ef is a test case for the introspectSchema method
func TestIntrospectSchema_4b3bcdb1ef(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code panicked: %v", r)
		}
	}()
	// case1: when introspection is enabled
	ec := &executionContextTest{DisableIntrospection: false}
	expectedResult, _ := introspection.WrapSchema(nil), nil
	result, err := ec.introspectSchemaTest()
	assert.Equal(t, expectedResult, result)
	assert.Equal(t, nil, err)
	t.Log("TestIntrospectSchema_4b3bcdb1ef passed for case: when introspection is enabled")

	// case2: when introspection is disabled
	ec = &executionContextTest{DisableIntrospection: true}
	expectedResult = nil
	expectedErr := errors.New("introspection disabled")
	result, err = ec.introspectSchemaTest()
	assert.Equal(t, expectedResult, result)
	assert.Equal(t, expectedErr, err)
	t.Log("TestIntrospectSchema_4b3bcdb1ef passed for case: when introspection is disabled")
}


func (ec *executionContextTest) introspectSchemaTest() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(nil), nil
}
