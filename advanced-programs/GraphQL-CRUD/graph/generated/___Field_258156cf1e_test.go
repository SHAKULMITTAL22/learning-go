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

func Test___Field_258156cf1e(t *testing.T) {
	ec := executionContext{
		OperationContext: &graphql.OperationContext{
			Variables: map[string]interface{}{},
		},
	}

	sel := ast.SelectionSet{
		Fields: []*ast.Field{
			{
				Name: "__typename",
			},
			{
				Name: "name",
			},
			{
				Name: "description",
			},
			{
				Name: "args",
			},
			{
				Name: "type",
			},
			{
				Name: "isDeprecated",
			},
			{
				Name: "deprecationReason",
			},
		},
	}

	obj := &introspection.Field{
		Name: "TestField",
		Type: &introspection.Type{
			Name: "TestType",
		},
		Args: []*introspection.InputValue{
			{
				Name: "testArg",
				Type: &introspection.Type{
					Name: "String",
				},
			},
		},
	}

	out := ec.___Field(context.Background(), sel, obj)

	if out == nil {
		t.Error("out is nil")
		return
	}

	if out.MarshalGQL() != `{"__typename":"__Field","name":"TestField","description":null,"args":[{"name":"testArg","type":"String"}],"type":"TestType","isDeprecated":false,"deprecationReason":null}` {
		t.Errorf("out.MarshalGQL() != `%s`", `{"__typename":"__Field","name":"TestField","description":null,"args":[{"name":"testArg","type":"String"}],"type":"TestType","isDeprecated":false,"deprecationReason":null}`)
	}
}

func Test___Field_Error(t *testing.T) {
	ec := executionContext{
		OperationContext: &graphql.OperationContext{
			Variables: map[string]interface{}{},
		},
	}

	sel := ast.SelectionSet{
		Fields: []*ast.Field{
			{
				Name: "__typename",
			},
			{
				Name: "name",
			},
			{
				Name: "description",
			},
			{
				Name: "args",
			},
			{
				Name: "type",
			},
			{
				Name: "isDeprecated",
			},
			{
				Name: "deprecationReason",
			},
		},
	}

	obj := &introspection.Field{
		Name: "TestField",
		Type: &introspection.Type{
			Name: "TestType",
		},
		Args: []*introspection.InputValue{
			{
				Name: "testArg",
				Type: &introspection.Type{
					Name: "String",
				},
			},
		},
	}

	ec.___Field_name = func(ctx context.Context, field *ast.Field, obj *introspection.Field) graphql.Marshaler {
		return graphql.Null
	}

	out := ec.___Field(context.Background(), sel, obj)

	if out != nil {
		t.Error("out is not nil")
		return
	}
}
