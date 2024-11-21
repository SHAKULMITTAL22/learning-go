package generated

import (
	"bytes"
	"context"
	"testing"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/graph/model"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
	"sync"
	"sync/atomic"
)

func TestMake_85494ce5df(t *testing.T) {
	var testexecutionContext executionContext
	selSet := ast.SelectionSet{}
	var testItems = []*model.Item{{ID: "1", Text: "Item1", Done: true}, {ID: "2", Text: "Item2", Done: false}}

    t.Run("Success case - multiple items", func(t *testing.T) {
		result := testexecutionContext.marshalNItem2ᚕᚖgithubᚗcomᚋtannergabrielᚋlearningᚑgoᚋadvancedᚑprogramsᚋGraphQLᚑCRUDᚋgraphᚋmodelᚐItem(context.Background(), selSet, testItems)
		if result == nil {
			t.Error("Expected non-nil result, got nil")
		} else {
			t.Log("Success case - multiple items passed")
		}
	})

	t.Run("Success case - single item", func(t *testing.T) {
		result := testexecutionContext.marshalNItem2ᚕᚖgithubᚗcomᚋtannergabrielᚋlearningᚑgoᚋadvancedᚑprogramsᚋGraphQLᚑCRUDᚋgraphᚋmodelᚐItem(context.Background(), selSet, testItems[:1])
		if result == nil {
			t.Error("Expected non-nil result, got nil")
		} else {
			t.Log("Success case - single item passed")
		}
	})

	t.Run("Failure case - empty item list", func(t *testing.T) {
		result := testexecutionContext.marshalNItem2ᚕᚖgithubᚗcomᚋtannergabrielᚋlearningᚑgoᚋadvancedᚑprogramsᚋGraphQLᚑCRUDᚋgraphᚋmodelᚐItem(context.Background(), selSet, []*model.Item{})
		if result != nil {
			t.Error("Expected nil, got non-nil result")
		} else {
			t.Log("Failure case - empty item list passed")
		}
	})
}
