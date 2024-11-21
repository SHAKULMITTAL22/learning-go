// Test generated by RoostGPT for test azure-32k-go using AI Type Azure Open AI and AI Model roost-gpt4-32k

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
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

func TestEnumValue_8076fad3d4(t *testing.T) {
	// Create an instance of executionContext 
	ec := &executionContext{}
	// Create an selection set
	set := ast.SelectionSet{}
	// Create a enumeration instance with isDeprecated: false
	enumValue := introspection.EnumValue{IsDeprecated: false}
	// Testing method with enumeration instance as isDeprecated: false
	t.Run("test1", func(t *testing.T) {
		if ec.marshalN__EnumValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValue(context.Background(), set, enumValue) == nil {
			t.Error("Expected the function to work with default EnumValue instance")
		} else {
			t.Log("Success case logged")
		}
	})

	// update enumeration instance with isDeprecated: true
	enumValue = introspection.EnumValue{IsDeprecated: true}
	// Testing method with enumeration instance as isDeprecated: true
	t.Run("test2", func(t *testing.T) {
		if ec.marshalN__EnumValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValue(context.Background(), set, enumValue) == nil {
			t.Error("Expected the function to work with EnumValue instance as deprecated")
		} else {
			t.Log("Success case logged")
		}
	})
}
