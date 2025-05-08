package graph

import (
	"context"
	"testing"

	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/database"
	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/graph/model"
)

type mutationResolver struct{}

func (r *mutationResolver) CreateItems(ctx context.Context, input model.NewItem) (*model.Item, error) {
	item := &model.Item{
		Title:  input.Title,
		Owner:  input.Owner,
		Rating: input.Rating,
	}
	// For simulation purpose
	return item, nil
}

func TestCreateItems(t *testing.T) {
	r := mutationResolver{}

	ctx := context.TODO()

	tests := []struct {
		name   string
		input  model.NewItem
		expect string
	}{
		{
			name:   "Test for Success",
			input:  model.NewItem{Title: "TestItem", Owner: "TestOwner", Rating: 5},
			expect: "TestItem",
		},
		{
			name:   "Test for No Item",
			input:  model.NewItem{},
			expect: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := r.CreateItems(ctx, tt.input)
			if gotErr != nil {
				t.Errorf("unexpected err: %v", gotErr)
				return
			}
			if got.Title != tt.expect {
				t.Errorf("expected %v, but got %v", tt.expect, got.Title)
			}
		})
	}
}
