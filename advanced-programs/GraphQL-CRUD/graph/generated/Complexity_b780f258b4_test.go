package generated

import (
	"context"
	"testing"
	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/graph/model"
)

// Define ComplexityRoot, ComplexityItem and ComplexityMutation struct
type ComplexityRoot struct {
	Item     *ComplexityItem
	Mutation *ComplexityMutation
}

type ComplexityItem struct {
	ID func(childComplexity int, args map[string]interface{}) int
}

type ComplexityMutation struct {
	CreateItems func(childComplexity int) int
}

// Define the executableSchema struct
type executableSchema struct {
	complexity *ComplexityRoot
}

func (e *executableSchema) Complexity(typeName, fieldName string, childComplexity int, args map[string]interface{}) (complexity int, ok bool) {
	switch typeName {
	case "Item":
		if e.complexity.Item != nil {
			switch fieldName {
			case "id":
				if e.complexity.Item.ID != nil {
					return e.complexity.Item.ID(childComplexity, args), true
				}
			}
		}
	case "Mutation":
		if e.complexity.Mutation != nil {
			switch fieldName {
			case "createItems":
				if e.complexity.Mutation.CreateItems != nil {
					return e.complexity.Mutation.CreateItems(childComplexity), true
				}
			}
		}
	}
	return 0, false
}

// TestComplexity_b780f258b4 is the test function for executableSchema's Complexity method.
func TestComplexity_b780f258b4(t *testing.T) {
	// Define the test cases ...
}
