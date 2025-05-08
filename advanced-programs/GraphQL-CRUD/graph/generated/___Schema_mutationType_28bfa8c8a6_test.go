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

// Create struct mock implementation of Introspection Middleware
type mockIntrospectionMiddleware struct {
	err error
}

func (m mockIntrospectionMiddleware) MutationType() *introspection.Type {
    if m.err != nil {
        return nil
    }
	return &introspection.Type{
        Name: "Type1",
    }
}

func Test___Schema_mutationType_28bfa8c8a6(t *testing.T) {
	tests := []struct {
		name           string
		obj            *mockIntrospectionMiddleware
		expectError    bool
		expectedResult *introspection.Type
	}{
		{
			name:           "Successful Response",
			obj:            &mockIntrospectionMiddleware{nil},
			expectError:    false,
			expectedResult: &introspection.Type{Name: "Type1"},
		},
		{
			name:        "Error Response",
			obj:         &mockIntrospectionMiddleware{errors.New("Error")},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ec := executionContext{}
			field := graphql.Field{
                Name: "fieldName",
            }
			res := ec.___Schema_mutationType(context.Background(), graphql.CollectedField{
                Field: &field,
            }, &introspection.Schema{MutationType: tt.obj})

			if tt.expectError {
				if _, ok := res.(*graphql.Null); !ok {
					t.Errorf("___Schema_mutationType(%v): expected error, got %v", tt.obj, res)
				}
			} else {
				if !reflect.DeepEqual(res, tt.expectedResult) {
					t.Errorf("___Schema_mutationType(%v): expected %v, got %v", tt.obj, tt.expectedResult, res)
				}
			}
		})
	}
}
