package generated

import (
	"context"
	"testing"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
)

type executionContext struct {}

// Mock ExecutionContext
func createMockExecutionContext() *executionContext {
	ec := executionContext{}
	// TODO: Initialize executionContext properties if needed
	return &ec
}

func (ec *executionContext) ___Directive_name(ctx context.Context, field graphql.CollectedField, dir *introspection.Directive) interface{} {
	// TODO: Implement the function body. For mock-up, return directive name.
	if dir != nil {
		return dir.Name
	} else {
		return nil
	}
}

// Mock CollectedField
func createMockCollectedField() graphql.CollectedField {
	cf := graphql.CollectedField{}
	// TODO: Initialize CollectedField properties if needed
	return cf
}

// Mock Directive
func createMockDirective() *introspection.Directive {
	return &introspection.Directive{Name: "TestDirective"}
}

func Test___Directive_name_ca4c018cf8(t *testing.T) {
	ctx := context.TODO()

	// Test case for successful execution
	{
		mockEc := createMockExecutionContext()
		mockField := createMockCollectedField()
		mockDirective := createMockDirective()
		expectedResult := "TestDirective"

		result := mockEc.___Directive_name(ctx, mockField, mockDirective)
		if strResult, ok := result.(string); ok {
			if strResult != expectedResult {
				t.Errorf("Failed to match expected result. Expected=%s, Got=%s", expectedResult, strResult)
			} else {
				t.Logf("Success: Expected result matched. Expected=%s, Got=%s", expectedResult, strResult)
			}
		} else {
			t.Errorf("Failed: Result is not string. Result=%v", result)
		}
	}

	// Test case for null directive
	{
		mockEc := createMockExecutionContext()
		mockField := createMockCollectedField()
		mockDirective := &introspection.Directive{}

		result := mockEc.___Directive_name(ctx, mockField, mockDirective)
		if strResult, ok := result.(string); ok {
			t.Errorf("Failed: Expected null as directive name is null, but got string. Result=%s", strResult)
		} else {
			t.Logf("Success: Expected null as directive name is null, and got null.")
		}
	}
}
