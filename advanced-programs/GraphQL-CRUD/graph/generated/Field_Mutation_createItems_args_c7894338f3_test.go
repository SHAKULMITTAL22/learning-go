package tests // Change `tests` to the appropriate package name where this test resides

import (
	"context"
	"testing"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stretchr/testify/assert"
	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/graph/model"
)

func TestField_Mutation_createItems_args_c7894338f3(t *testing.T) {
	ec := NewExecutionContext(graphql.OperationDefinition{}, graphql.RandomizedInput(nil, nil))

	t.Run("Test the successful mutation for creation of items", func(t *testing.T) {
		input := model.NewItem{Name:"Item1", Description:"This is a sample item"}
		args := map[string]interface{}{
			"input":input,
		}
		_, err := ec.field_Mutation_createItems_args(context.Background(), args)
		assert.Nil(t, err, "Expected Nil, but got ", err)
		t.Log("Best case, pass!! arguments were: ", args)
	})

	t.Run("Test the failed mutation for creation of items", func(t *testing.T) {
		input := "Item1"
		args := map[string]interface{}{
			"input":input,
		}
		_, err := ec.field_Mutation_createItems_args(context.Background(), args)
		assert.Error(t, err, "Expected an error, but didn't get one")
		t.Log("Error case, pass!! arguments were: ", args)
	})
}
