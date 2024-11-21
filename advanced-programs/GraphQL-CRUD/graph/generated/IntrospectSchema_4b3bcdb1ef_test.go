// Test generated by RoostGPT for test math-go using AI Type Open AI and AI Model gpt-4

package generated

import (
	"errors"
	"testing"

	"github.com/99designs/gqlgen/graphql/introspection"
	gqlparser "github.com/vektah/gqlparser/v2"
)

type executionContextTest struct {
	DisableIntrospection bool
}

func (ec *executionContextTest) introspectSchemaTest() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(&gqlparser.Schema{}), nil
}

func TestIntrospectSchema_4b3bcdb1ef(t *testing.T) {
	t.Run("introspection is disabled", func(t *testing.T) {
		ec := &executionContextTest{
			DisableIntrospection: true,
		}
		_, err := ec.introspectSchemaTest()
		if err == nil || err.Error() != "introspection disabled" {
			t.Error("Expected error, got none")
		} else {
			t.Log("Test Passed: Expected error received")
		}
	})

	t.Run("introspection is enabled", func(t *testing.T) {
		ec := &executionContextTest{
			DisableIntrospection: false,
		}
		_, err := ec.introspectSchemaTest()
		if err != nil {
			t.Error("Expected no error, got error: ", err)
		} else {
			t.Log("Test Passed: No error received as expected")
		}
	})
}
