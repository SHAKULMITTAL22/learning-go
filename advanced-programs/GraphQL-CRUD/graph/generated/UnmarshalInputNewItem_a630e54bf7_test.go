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

func TestUnmarshalInputNewItem_a630e54bf7(t *testing.T) {
	ec := executionContext{unmarshalNString2string: func(ctx context.Context, v interface{}) (string, error) {
		return "test", nil
	}, unmarshalNInt2int: func(ctx context.Context, v interface{}) (int, error) {
		return 1, nil
	}}
	obj := map[string]interface{}{
		"title":  "test",
		"owner":  "test",
		"rating": 1,
	}

	item, err := ec.unmarshalInputNewItem(context.Background(), obj)
	if err != nil {
		t.Error(err)
	}

	if item.Title != "test" {
		t.Error("expected title to be 'test'")
	}
	if item.Owner != "test" {
		t.Error("expected owner to be 'test'")
	}
	if item.Rating != 1 {
		t.Error("expected rating to be 1")
	}
}

func TestUnmarshalInputNewItem_error(t *testing.T) {
	ec := executionContext{unmarshalNString2string: func(ctx context.Context, v interface{}) (string, error) {
		return "", errors.New("error")
	}, unmarshalNInt2int: func(ctx context.Context, v interface{}) (int, error) {
		return 0, errors.New("error")
	}}
	obj := map[string]interface{}{
		"title":  "test",
		"owner":  "test",
		"rating": 1,
	}

	_, err := ec.unmarshalInputNewItem(context.Background(), obj)
	if err == nil {
		t.Error("expected error")
	}
}
