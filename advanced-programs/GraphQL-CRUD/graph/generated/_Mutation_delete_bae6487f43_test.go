package generated

import (
	"context"
	"errors"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	model "github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/graph/model"
)

type mockResolver struct{}

func (r *mockResolver) Mutation() model.MutationResolver {
	return &mockMutationResolver{}
}

type mockMutationResolver struct{}

func (r *mockMutationResolver) Delete(ctx context.Context, id int) (string, error) {
	if id == 1 {
		return "Deleted", nil
	}
	return "", errors.New("Failed to delete")
}

func Test_Mutation_delete(t *testing.T) {
	resolver := &mockResolver{}
	tests := []struct {
		name    string
		id      int
		want    string
		wantErr bool
	}{
		{
			name:    "Test Case 1: Successful Deletion",
			id:      1,
			want:    "Deleted",
			wantErr: false,
		},
		{
			name:    "Test Case 2: Failed Deletion",
			id:      2,
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			got, err := resolver.Mutation().Delete(ctx, tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
