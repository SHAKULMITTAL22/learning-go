package generated

import (
	"context"
	"testing"

	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/graph/model"
)

type ComplexityRoot struct {
	Item     *ComplexityItem
	Mutation *ComplexityMutation
	Query    *ComplexityQuery
}

type ComplexityItem struct {
	ID func(childComplexity int) int
}

type ComplexityMutation struct {
	CreateItems func(childComplexity int, input *model.NewItem) int
	Delete      func(childComplexity int, id int) int
}

type ComplexityQuery struct {
	Item  func(childComplexity int, id int) int
	Items func(childComplexity int) int
}

type executableSchema struct {
	complexity *ComplexityRoot
}

func (e *executableSchema) Complexity(typeName string, fieldName string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	// Implement your complexity calculation logic here
	return 0, false
}

func TestComplexity_b780f258b4(t *testing.T) {
	// Test case 1: testing for "Item.id"
	t.Run("Item.id", func(t *testing.T) {
		e := &executableSchema{
			complexity: &ComplexityRoot{
				Item: &ComplexityItem{
					ID: func(childComplexity int) int {
						return childComplexity + 1
					},
				},
			},
		}
		complexity, ok := e.Complexity("Item", "id", 1, nil)
		if !ok || complexity != 2 {
			t.Error("Expected complexity=2, got", complexity)
		} else {
			t.Log("Test case passed successfully")
		}
	})

	// Test case 2: testing for "Mutation.createItems"
	t.Run("Mutation.createItems", func(t *testing.T) {
		e := &executableSchema{
			complexity: &ComplexityRoot{
				Mutation: &ComplexityMutation{
					CreateItems: func(childComplexity int, input *model.NewItem) int {
						return childComplexity + 1
					},
				},
			},
		}
		complexity, ok := e.Complexity("Mutation", "createItems", 1, map[string]interface{}{
			"input": &model.NewItem{},
		})
		if !ok || complexity != 2 {
			t.Error("Expected complexity=2, got", complexity)
		} else {
			t.Log("Test case passed successfully")
		}
	})

	// Test case 3: testing for "Mutation.delete"
	t.Run("Mutation.delete", func(t *testing.T) {
		e := &executableSchema{
			complexity: &ComplexityRoot{
				Mutation: &ComplexityMutation{
					Delete: func(childComplexity int, id int) int {
						return childComplexity + 1
					},
				},
			},
		}
		complexity, ok := e.Complexity("Mutation", "delete", 1, map[string]interface{}{
			"id": 1,
		})
		if !ok || complexity != 2 {
			t.Error("Expected complexity=2, got", complexity)
		} else {
			t.Log("Test case passed successfully")
		}
	})

	// Test case 4: testing for "Query.item"
	t.Run("Query.item", func(t *testing.T) {
		e := &executableSchema{
			complexity: &ComplexityRoot{
				Query: &ComplexityQuery{
					Item: func(childComplexity int, id int) int {
						return childComplexity + 1
					},
				},
			},
		}
		complexity, ok := e.Complexity("Query", "item", 1, map[string]interface{}{
			"id": 1,
		})
		if !ok || complexity != 2 {
			t.Error("Expected complexity=2, got", complexity)
		} else {
			t.Log("Test case passed successfully")
		}
	})

	// Test case 5: testing for "Query.items"
	t.Run("Query.items", func(t *testing.T) {
		e := &executableSchema{
			complexity: &ComplexityRoot{
				Query: &ComplexityQuery{
					Items: func(childComplexity int) int {
						return childComplexity + 1
					},
				},
			},
		}
		complexity, ok := e.Complexity("Query", "items", 1, nil)
		if !ok || complexity != 2 {
			t.Error("Expected complexity=2, got", complexity)
		} else {
			t.Log("Test case passed successfully")
		}
	})
}
