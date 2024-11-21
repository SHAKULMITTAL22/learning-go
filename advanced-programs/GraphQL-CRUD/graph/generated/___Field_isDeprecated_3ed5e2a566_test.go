// Test generated by RoostGPT for test math-go using AI Type Open AI and AI Model gpt-4

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
	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/graph/model"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

func Test___Field_isDeprecated_3ed5e2a566(t *testing.T) {
	ec := &executionContext{}
	field := graphql.CollectedField{}
	obj := &introspection.Field{}

	// Test case 1: Successful execution
	t.Run("Test case 1: Successful execution", func(t *testing.T) {
		obj.DeprecationReason = "Test deprecation"
		result := ec.___Field_isDeprecated(context.Background(), field, obj)
		if result == graphql.Null {
			t.Error("Test case 1 failed: expected non null result")
		} else {
			t.Log("Test case 1 passed")
		}
	})

	// Test case 2: Error execution
	t.Run("Test case 2: Error execution", func(t *testing.T) {
		obj = nil
		defer func() {
			if r := recover(); r != nil {
				t.Log("Test case 2 passed")
			} else {
				t.Error("Test case 2 failed: expected panic")
			}
		}()
		ec.___Field_isDeprecated(context.Background(), field, obj)
	})
}
