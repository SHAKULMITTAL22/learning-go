package generated

import (
	"bytes"
	"context"
	"errors"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/vektah/gqlparser/v2/ast"
)

type  executionContext struct{}

func (ec *executionContext) ___Type_possibleTypes(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) []*introspection.Type{
	return obj.PossibleTypes
}

func Test___Type_possibleTypes(t *testing.T) {
	var dummyExecutionContext = &executionContext{}
	ctx := context.Background()

	t.Run("Test with nil introspection type", func(t *testing.T) {
		field := graphql.CollectedField{}

		obj := &introspection.Type{}
		res := dummyExecutionContext.___Type_possibleTypes(ctx, field, obj)

		if res == nil {
			t.Errorf("Received nil result")
		}
	})

	t.Run("Test with a valid introspection type ", func(t *testing.T) {
		field := graphql.CollectedField{}

		//Initialize the introspection Type
		obj := &introspection.Type{Name: "your_type_name",
			Description: "your_type_description",
			PossibleTypes: []*introspection.Type{
				{Name: "sub_type_1", Description: "desc_1"},
				{Name: "sub_type_2", Description: "desc_2"},
			},
		}
		res := dummyExecutionContext.___Type_possibleTypes(ctx, field, obj)

		if res == nil {
			t.Errorf("Received nil result")
		}
	})
}
