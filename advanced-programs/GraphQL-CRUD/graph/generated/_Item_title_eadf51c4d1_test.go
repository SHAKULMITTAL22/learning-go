package generated

import (
	"context"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/graph/model"
)

func Test_Item_title_eadf51c4d1(t *testing.T) {
	ec := &executionContext{}
	ctx := context.Background()
	field := graphql.CollectedField{}
	obj := &model.Item{Title: "Test Item"}

	t.Run("Successful Item Title", func(t *testing.T) {
		res := ec._Item_title(ctx, field, obj)
		if res == nil {
			t.Error("Expected non-nil result, but got nil")
		}
		t.Log("Test Successful Item Title passed")
	})

	t.Run("Item Title with nil object", func(t *testing.T) {
		res := ec._Item_title(ctx, field, nil)
		if res != graphql.Null {
			t.Error("Expected graphql.Null, but got non-null result")
		}
		t.Log("Test Item Title with nil object passed")
	})

	t.Run("Item Title with empty title", func(t *testing.T) {
		obj := &model.Item{Title: ""}
		res := ec._Item_title(ctx, field, obj)
		if res != graphql.Null {
			t.Error("Expected graphql.Null, but got non-null result")
		}
		t.Log("Test Item Title with empty title passed")
	})
}
