package generated

import (
	"bytes"
	"context"
	"errors"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/graph/model"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

func TestNewItem_97d2d40b62(t *testing.T) {
	ec := executionContext{
		unmarshalInputNewItem: func(ctx context.Context, v interface{}) (model.NewItem, error) {
			return model.NewItem{}, nil
		},
	}

	// Success case
	v := `{"name": "test", "description": "test"}`
	_, err := ec.unmarshalNNewItem2githubᚗcomᚋtannergabrielᚋlearningᚑgoᚋadvancedᚑprogramsᚋGraphQLᚑCRUDᚋgraphᚋmodelᚐNewItem(context.Background(), v)
	if err != nil {
		t.Error(err)
	}

	// Error case
	v = `{"name": 123, "description": "test"}`
	_, err = ec.unmarshalNNewItem2githubᚗcomᚋtannergabrielᚋlearningᚑgoᚋadvancedᚑprogramsᚋGraphQLᚑCRUDᚋgraphᚋmodelᚐNewItem(context.Background(), v)
	if err == nil {
		t.Error("expected error")
	}
}
