package generated

import (
	"context"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/graph/model"
	"strconv"
)

type mockResolver struct{}

func (m *mockResolver) Query() QueryResolver {
	return &mockQueryResolver{}
}

type mockQueryResolver struct{}

func (m *mockQueryResolver) Item(ctx context.Context, id int) (*model.Item, error) {
	if id == 0 {
		return nil, nil
	}
	return &model.Item{ID: strconv.Itoa(id)}, nil
}

func Test_Query_item_f90ef048de(t *testing.T) {
	t.Run("Success case", func(t *testing.T) {
		ec := &executionContext{
			resolvers: &mockResolver{},
		}
		field := graphql.CollectedField{}
		result := ec._Query_item(context.Background(), field)
		if result == graphql.Null {
			t.Error("Expected result not to be null")
		} else {
			t.Log("Success case passed")
		}
	})

	t.Run("Failure case", func(t *testing.T) {
		ec := &executionContext{
			resolvers: &mockResolver{},
		}
		field := graphql.CollectedField{}
		result := ec._Query_item(context.Background(), field)
		if result != graphql.Null {
			t.Error("Expected result to be null")
		} else {
			t.Log("Failure case passed")
		}
	})
}
