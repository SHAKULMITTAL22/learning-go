package generated

import (
	"testing"
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/graph/model"
)

type MockResolver struct {
	mock.Mock
}

func (m *MockResolver) Items(ctx context.Context) ([]*model.Item, error) {
	ret := m.Called(ctx)

	var r0 []*model.Item
	if rf, ok := ret.Get(0).(func(context.Context) []*model.Item); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Item)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type ExecutionContext struct {
	Resolver *MockResolver
}

func (ec *ExecutionContext) QueryItems(ctx context.Context, cf graphql.CollectedField) []*model.Item {
	items, _ := ec.Resolver.Items(ctx)
	return items
}

func TestQueryItems(t *testing.T) {
	mockResolver := new(MockResolver)
	ec := &ExecutionContext{
		Resolver: mockResolver,
	}

	items := []*model.Item{
		{ID: "1", Title: "Item 1", Price: 100},
		{ID: "2", Title: "Item 2", Price: 200},
	}

	t.Run("Items Success", func(t *testing.T) {
		mockResolver.On("Items", mock.Anything).Return(items, nil).Once()

		result := ec.QueryItems(context.Background(), graphql.CollectedField{})
		assert.NotNil(t, result)

		mockResolver.AssertExpectations(t)
	})

	t.Run("Items Error", func(t *testing.T) {
		mockResolver.On("Items", mock.Anything).Return(nil, assert.AnError).Once()

		result := ec.QueryItems(context.Background(), graphql.CollectedField{})
		assert.Nil(t, result)

		mockResolver.AssertExpectations(t)
	})

	t.Run("Items Nil", func(t *testing.T) {
		mockResolver.On("Items", mock.Anything).Return(nil, nil).Once()

		result := ec.QueryItems(context.Background(), graphql.CollectedField{})
		assert.Nil(t, result)

		mockResolver.AssertExpectations(t)
	})
}
