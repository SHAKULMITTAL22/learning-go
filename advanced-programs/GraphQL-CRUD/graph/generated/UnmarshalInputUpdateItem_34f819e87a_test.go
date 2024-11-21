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

func TestUnmarshalInputUpdateItem_34f819e87a(t *testing.T) {
	ec := executionContext{
		unmarshalInputUpdateItem: func(ctx context.Context, obj interface{}) (model.UpdateItem, error) {
			var it model.UpdateItem
			var asMap = obj.(map[string]interface{})

			for k, v := range asMap {
				switch k {
				case "title":
					var err error
					it.Title, err = ec.unmarshalOString2ᚖstring(ctx, v)
					if err != nil {
						return it, err
					}
				case "owner":
					var err error
					it.Owner, err = ec.unmarshalOString2ᚖstring(ctx, v)
					if err != nil {
						return it, err
					}
				case "rating":
					var err error
					it.Rating, err = ec.unmarshalOInt2ᚖint(ctx, v)
					if err != nil {
						return it, err
					}
				}
			}

			return it, nil
		},
	}

	// Success case
	obj := map[string]interface{}{
		"title":  "New Title",
		"owner":  "New Owner",
		"rating": 5,
	}
	expected := model.UpdateItem{
		Title:  "New Title",
		Owner:  "New Owner",
		Rating: 5,
	}
	actual, err := ec.unmarshalInputUpdateItem(context.Background(), obj)
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	// Error case: missing required field
	obj = map[string]interface{}{
		"title": "New Title",
	}
	_, err = ec.unmarshalInputUpdateItem(context.Background(), obj)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if !errors.Is(err, model.ErrMissingRequiredField) {
		t.Errorf("Expected error %v, got %v", model.ErrMissingRequiredField, err)
	}

	// Error case: invalid field type
	obj = map[string]interface{}{
		"title":  "New Title",
		"owner":  123, // Invalid type for owner field
		"rating": 5,
	}
	_, err = ec.unmarshalInputUpdateItem(context.Background(), obj)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if !errors.Is(err, model.ErrInvalidFieldType) {
		t.Errorf("Expected error %v, got %v", model.ErrInvalidFieldType, err)
	}
}
