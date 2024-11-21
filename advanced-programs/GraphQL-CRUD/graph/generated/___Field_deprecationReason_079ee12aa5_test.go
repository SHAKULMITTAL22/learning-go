// Test case regenerated by RoostGPT for test azure-32k-go using AI Type Azure Open AI and AI Model roost-gpt4-32k

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
	"testing"
)

// Test___Field_deprecationReason_079ee12aa5 tests the ___Field_deprecationReason method
func Test___Field_deprecationReason_079ee12aa5(t *testing.T) {
	t.Run("Test with a deprecated field", func(t *testing.T) {
		ctx := context.Background()
		ec := &executionContext{}
		field := graphql.CollectedField{}
		obj := &introspection.Field{DeprecationReason: func() *string {
			str := "Deprecated"
			return &str
		}}
		result := ec.___Field_deprecationReason(ctx, field, obj)
		if result == nil {
			t.Error("Expected marshaled string, but got nil")
		} else {
			t.Log("Test passed with deprecated field")
		}
	})
	
	t.Run("Test with a non-deprecated field", func(t *testing.T) {
		ctx := context.Background()
		ec := &executionContext{}
		field := graphql.CollectedField{}
		obj := &introspection.Field{}
		result := ec.___Field_deprecationReason(ctx, field, obj)
		if result != nil {
			t.Error("Expected null, but got a value")
		} else {
			t.Log("Test passed with non-deprecated field")
		}
	})
	
	t.Run("Test with panicking function", func(t *testing.T) {
		ctx := context.Background()
		ec := &executionContext{}
		field := graphql.CollectedField{}
		obj := &introspection.Field{DeprecationReason: func() *string {
			panic("I'm panicking!")
		}}
		result := ec.___Field_deprecationReason(ctx, field, obj)
		if result != nil {
			t.Error("Expected null, but got a value")
		} else {
			t.Log("Test passed with panicking function")
		}
	})
}
