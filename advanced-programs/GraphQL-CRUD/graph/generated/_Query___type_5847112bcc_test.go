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
)

func Test_Query___type_5847112bcc(t *testing.T) {
	ec := executionContext{
		introspectType: func(name string) *introspection.Type {
			return &introspection.Type{
				Name: name,
			}
		},
	}

	// Success case
	t.Run("success", func(t *testing.T) {
		args := map[string]interface{}{
			"name": "Foo",
		}
		ctx := context.Background()
		res := ec._Query___type(ctx, graphql.CollectedField{
			ArgumentMap: func(variables graphql.Variables) map[string]interface{} {
				return args
			},
		})
		if res == nil {
			t.Error("expected non-nil result")
		}
	})

	// Error case
	t.Run("error", func(t *testing.T) {
		args := map[string]interface{}{
			"name": "Bar",
		}
		ctx := context.Background()
		res := ec._Query___type(ctx, graphql.CollectedField{
			ArgumentMap: func(variables graphql.Variables) map[string]interface{} {
				return args
			},
		})
		if res != nil {
			t.Error("expected nil result")
		}
	})
}
