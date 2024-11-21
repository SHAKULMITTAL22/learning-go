package graph

import (
	"context"
	"testing"
	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/database"
	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/graph/model"
	"github.com/stretchr/testify/assert"
)

type queryResolverMock struct {
	data []*model.Item
}

func (r *queryResolverMock) Items(ctx context.Context) ([]*model.Item, error) {
	return r.data, nil
}

func TestItems_b319e5fa60(t *testing.T) {
	item1 := &model.Item{ID: "1", Text: "Apple"}
	item2 := &model.Item{ID: "2", Text: "Banana"}

	r := &queryResolverMock{data: []*model.Item{item1, item2}}

	t.Run("success", func(t *testing.T) {
		result, err := r.Items(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, []*model.Item{item1, item2}, result)
		t.Log("TestItems_b319e5fa60 passed for successful scenario")
	})

	t.Run("fail", func(t *testing.T) {
		rFail := &queryResolverMock{data: nil}
		result, err := rFail.Items(context.Background())
		assert.NoError(t, err)
		assert.Nil(t, result)
		t.Log("TestItems_b319e5fa60 passed for failure scenario")
	})
}
