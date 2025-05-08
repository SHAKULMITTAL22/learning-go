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

func TestItem_1c96e05301(t *testing.T) {
	ec := executionContext{
		introspection.Schema{},
		map[string]graphql.Marshaler{},
		map[string]graphql.Type{},
		map[string]graphql.Field{},
		bytes.NewBuffer(nil),
		context.Background(),
		&sync.Mutex{},
		&atomic.Value{},
	}

	// Success case
	v := &model.Item{
		ID:          "1",
		Name:        "Item 1",
		Description: "This is item 1",
		Price:       10.0,
	}
	marshaler := ec.marshalNItem2ᚖgithubᚗcomᚋtannergabrielᚋlearningᚑgoᚋadvancedᚑprogramsᚋGraphQLᚑCRUDᚋgraphᚋmodelᚐItem(context.Background(), nil, v)
	if _, ok := marshaler.(*graphql.NullValue); !ok {
		t.Errorf("Expected *graphql.NullValue, got %T", marshaler)
	}

	// Failure case
	v = nil
	marshaler = ec.marshalNItem2ᚖgithubᚗcomᚋtannergabrielᚋlearningᚑgoᚋadvancedᚑprogramsᚋGraphQLᚑCRUDᚋgraphᚋmodelᚐItem(context.Background(), nil, v)
	if _, ok := marshaler.(*graphql.NullValue); ok {
		t.Errorf("Expected non-null value, got *graphql.NullValue")
	}
}
