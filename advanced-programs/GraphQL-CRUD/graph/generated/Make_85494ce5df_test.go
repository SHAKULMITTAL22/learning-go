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

func TestMake_85494ce5df(t *testing.T) {
	ec := executionContext{
		introspection: introspection.NewSchema(gqlparser.MustLoadSchema(&ast.Source{Name: "schema.graphql"})),
	}

	// Success case
	v := []*model.Item{
		{
			ID:          "1",
			Name:        "Item 1",
			Description: "Description 1",
			Price:       10.0,
		},
		{
			ID:          "2",
			Name:        "Item 2",
			Description: "Description 2",
			Price:       20.0,
		},
	}
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
	ret := ec.marshalNItem2ᚕᚖgithubᚗcomᚋtannergabrielᚋlearningᚑgoᚋadvancedᚑprogramsᚋGraphQLᚑCRUDᚋgraphᚋmodelᚐItemᚄ(context.Background(), sel, v)
	if len(ret) != len(v) {
		t.Errorf("Expected %d items, got %d", len(v), len(ret))
	}
	for i, item := range ret {
		if item == nil {
			t.Errorf("Item %d is nil", i)
		} else if item.ID != v[i].ID {
			t.Errorf("Expected item %d to have ID %s, got %s", i, v[i].ID, item.ID)
		} else if item.Name != v[i].Name {
			t.Errorf("Expected item %d to have Name %s, got %s", i, v[i].Name, item.Name)
		} else if item.Description != v[i].Description {
			t.Errorf("Expected item %d to have Description %s, got %s", i, v[i].Description, item.Description)
		} else if item.Price != v[i].Price {
			t.Errorf("Expected item %d to have Price %f, got %f", i, v[i].Price, item.Price)
		}
	}

	// Error case
	v = []*model.Item{
		{
			ID:          "1",
			Name:        "Item 1",
			Description: "Description 1",
			Price:       10.0,
		},
		{
			ID:          "2",
			Name:        "Item 2",
			Description: "Description 2",
			Price:       20.0,
		},
		{
			ID:          "3",
			Name:        "Item 3",
			Description: "Description 3",
			Price:       30.0,
		},
	}
	sel = ast.SelectionSet{
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
			{
				Name: "invalidField",
			},
		},
	}
	ret = ec.marshalNItem2ᚕᚖgithubᚗcomᚋtannergabrielᚋlearningᚑgoᚋadvancedᚑprogramsᚋGraphQLᚑCRUDᚋgraphᚋmodelᚐItemᚄ(context.Background(), sel, v)
	if len(ret) != len(v) {
		t.Errorf("Expected %d items, got %d", len(v), len(ret))
	}
	for i, item := range ret {
		if item == nil {
			t.Errorf("Item %d is nil", i)
		} else if item.ID != v[i].ID {
			t.Errorf("Expected item %d to have ID %s, got %s", i, v[i].ID, item.ID)
		} else if item.Name != v[i].Name {
			t.Errorf("Expected item %d to have Name %s, got %s", i, v[i].Name, item.Name)
		} else if item.Description != v[i].Description {
			t.Errorf("Expected item %d to have Description %s, got %s", i, v[i].Description, item.Description)
		} else if item.Price != v[i].Price {
			t.Errorf("Expected item %d to have Price %f, got %f", i, v[i].Price, item.Price)
		}
	}
}
