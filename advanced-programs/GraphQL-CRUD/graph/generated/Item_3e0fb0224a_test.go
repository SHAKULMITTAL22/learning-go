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

func TestItem_3e0fb0224a(t *testing.T) {
	ec := executionContext{
		directives: make(map[string]func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error)),
		resolvers:  ec.defaultResolvers,
	}

	ctx := context.Background()
	sel := ast.SelectionSet{
		Fields: []*ast.Field{
			{
				Name: "id",
			},
			{
				Name: "name",
			},
			{
				Name: "description",
			},
			{
				Name: "price",
			},
		},
	}
	v := model.Item{
		ID:          "1",
		Name:        "Test Item",
		Description: "This is a test item",
		Price:       10.0,
	}

	buf := bytes.NewBufferString("")
	res := ec.marshalNItem2githubᚗcomᚋtannergabrielᚋlearningᚑgoᚋadvancedᚑprogramsᚋGraphQLᚑCRUDᚋgraphᚋmodelᚐItem(ctx, sel, v)
	if err := res.MarshalGQL(buf); err != nil {
		t.Error(err)
	}

	expected := `{"id":"1","name":"Test Item","description":"This is a test item","price":10}`
	if buf.String() != expected {
		t.Errorf("Expected %s, got %s", expected, buf.String())
	}
}

func TestItem_3e0fb0224a_Error(t *testing.T) {
	ec := executionContext{
		directives: make(map[string]func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error)),
		resolvers:  ec.defaultResolvers,
	}

	ctx := context.Background()
	sel := ast.SelectionSet{
		Fields: []*ast.Field{
			{
				Name: "id",
			},
			{
				Name: "name",
			},
			{
				Name: "description",
			},
			{
				Name: "price",
			},
		},
	}
	v := model.Item{
		ID:          "1",
		Name:        "Test Item",
		Description: "This is a test item",
		Price:       -10.0, // Invalid price
	}

	buf := bytes.NewBufferString("")
	res := ec.marshalNItem2githubᚗcomᚋtannergabrielᚋlearningᚑgoᚋadvancedᚑprogramsᚋGraphQLᚑCRUDᚋgraphᚋmodelᚐItem(ctx, sel, v)
	if err := res.MarshalGQL(buf); err == nil {
		t.Error("Expected error, got nil")
	} else if !errors.Is(err, model.ErrInvalidPrice) {
		t.Errorf("Expected error %v, got %v", model.ErrInvalidPrice, err)
	}
}
