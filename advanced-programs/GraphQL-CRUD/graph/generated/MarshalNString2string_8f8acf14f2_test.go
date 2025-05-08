package generated

import (
	"context"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

type executionContext struct{}

func (e *executionContext) marshalNString2string(ctx context.Context, sel ast.SelectionSet, str string) interface{} {
	return nil
}

func TestMarshalNString2string_8f8acf14f2(t *testing.T) {
	t.Parallel()
	ec := &executionContext{}
	ctx := context.Background()
	sel := make(ast.SelectionSet, 0)

	{
		v := "test_string"
		if res := ec.marshalNString2string(ctx, sel, v); res == graphql.Null {
			t.Errorf("marshalNString2string(%v, %v, %v): got = %v, want: non-null", ctx, sel, v, res)
		} else {
			t.Logf("success: marshalNString2string(%v, %v, %v): got = %v", ctx, sel, v, res)
		}
	}

	{
		// setting up a field error in the context
		err := graphql.AddError(ctx, &gqlerror.Error{
			Message: "Field error",
			Locations: []ast.Location{
				{Line: 1, Column: 1},
			},
			Path: ast.Path{
				ast.PathElement{
					Name: "test",
				},
			},
		})

		v := ""
		if res := ec.marshalNString2string(ctx, sel, v); err == nil && res != graphql.Null {
			t.Errorf("marshalNString2string(%v, %v, %v): got = %v, want: graphql.Null", ctx, sel, v, res)
		} else {
			t.Logf("success: marshalNString2string(%v, %v, %v): got = %v", ctx, sel, v, res)
		}
	}
}
