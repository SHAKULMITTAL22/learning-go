package generated

import (
	"context"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/vektah/gqlparser/v2/ast"
)

// Execution context
type ExecutionContext struct {
	Schema *introspection.Schema
}

// Creating a new ExecutionContext
func newExecutionContext(ctx context.Context, object *introspection.Schema, fields []*ast.Field) *ExecutionContext {
	return &ExecutionContext{
		Schema: object,
	}
}

// Empty Execution Context
type EmptyExecutionContext struct {
	Schema *introspection.Schema
}

// Creating a new EmptyExecutionContext
func newEmptyExecutionContext(ctx context.Context, object *introspection.Schema, fields []*ast.Field) *EmptyExecutionContext {
	return &EmptyExecutionContext{
		Schema: object,
	}
}

// Writing schema to context
func (ec *ExecutionContext) ___Schema(ctx context.Context, sel ast.SelectionSet, object *introspection.Schema) graphql.Marshaler {
	return graphql.Null{}
}

// Writing empty schema to context
func (ec *EmptyExecutionContext) ___Schema(ctx context.Context, sel ast.SelectionSet, object *introspection.Schema) graphql.Marshaler {
	return graphql.Null{}
}

func Test___Schema_335b8a3ee1(t *testing.T) {
	t.Run("Successful Marshalar return", func(t *testing.T) {
		ctx := context.TODO()
		object := &introspection.Schema{
			QueryType:        &introspection.TypeRef{Name: "queryType"},
			MutationType:     &introspection.TypeRef{Name: "mutationType"},
			SubscriptionType: &introspection.TypeRef{Name: "subscriptionType"},
			Types:            []introspection.Type{{Name: "types"}},
			Directives:       []introspection.Directive{{Name: "directives"}},
		}
		fields := []*ast.Field{{Name: "types"}, {Name: "queryType"}, {Name: "mutationType"}, {Name: "subscriptionType"}, {Name: "directives"}}
		sel := ast.SelectionSet{}
		for _, field := range fields {
			sel = append(sel, &ast.FieldSelection{Field: field})
		}
		ec := newExecutionContext(ctx, object, fields)

		result := ec.___Schema(ctx, sel, object)

		if _, ok := result.(graphql.Null); ok {
			t.Errorf("___Schema() = %v, want %v", result, "graphql.Marshaler")
		}

		t.Log("Test Successful Marshalar return passed")
	})

	t.Run("Returns graphql.Null ", func(t *testing.T) {
		ctx := context.TODO()
		object := &introspection.Schema{
			QueryType:        &introspection.TypeRef{Name: "queryType"},
			MutationType:     &introspection.TypeRef{Name: "mutationType"},
			SubscriptionType: &introspection.TypeRef{Name: "subscriptionType"},
			Types:            []introspection.Type{{Name: "types"}},
			Directives:       []introspection.Directive{{Name: "directives"}},
		}
		fields := []*ast.Field{{Name: "types"}, {Name: "queryType"}, {Name: "mutationType"}, {Name: "subscriptionType"}, {Name: "directives"}}
		sel := ast.SelectionSet{}
		for _, field := range fields {
			sel = append(sel, &ast.FieldSelection{Field: field})
		}
		ec := newEmptyExecutionContext(ctx, object, fields)

		result := ec.___Schema(ctx, sel, object)

		if _, ok := result.(graphql.Null); !ok {
			t.Errorf("___Schema() = %v, want %v", result, "graphql.Null")
		}

		t.Log("Test Returns graphql.Null passed")
	})

	t.Run("Method panic", func(t *testing.T) {
		ctx := context.TODO()
		object := &introspection.Schema{
			QueryType:        &introspection.TypeRef{Name: "queryType"},
			MutationType:     &introspection.TypeRef{Name: "mutationType"},
			SubscriptionType: &introspection.TypeRef{Name: "subscriptionType"},
			Types:            []introspection.Type{{Name: "types"}},
			Directives:       []introspection.Directive{{Name: "directives"}},
		}
		fields := []*ast.Field{{Name: "invalid"}}
		sel := ast.SelectionSet{}
		for _, field := range fields {
			sel = append(sel, &ast.FieldSelection{Field: field})
		}
		ec := newExecutionContext(ctx, object, fields)

		defer func() {
			if r := recover(); r == nil {
				t.Error("___Schema() should have panicked!")
			}
		}()

		ec.___Schema(ctx, sel, object)

		t.Log("Test Method panic passed")
	})
}
