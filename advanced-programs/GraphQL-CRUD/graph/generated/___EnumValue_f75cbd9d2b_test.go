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

func Test___EnumValue_f75cbd9d2b(t *testing.T) {
	ec := executionContext{
		OperationContext: &graphql.OperationContext{
			Variables: map[string]interface{}{},
		},
	}

	obj := &introspection.EnumValue{
		Name: "EnumValue",
		Description: "Description",
		IsDeprecated: true,
		DeprecationReason: "DeprecationReason",
	}

	sel := ast.SelectionSet{
		Selections: []ast.Selection{
			&ast.Field{
				Name: "__typename",
			},
			&ast.Field{
				Name: "name",
			},
			&ast.Field{
				Name: "description",
			},
			&ast.Field{
				Name: "isDeprecated",
			},
			&ast.Field{
				Name: "deprecationReason",
			},
		},
	}

	out := ec.___EnumValue(context.Background(), sel, obj)

	if out == nil {
		t.Error("out is nil")
		return
	}

	if out.MarshalGQL() != `{"__typename":"__EnumValue","name":"EnumValue","description":"Description","isDeprecated":true,"deprecationReason":"DeprecationReason"}` {
		t.Error("out.MarshalGQL() is not equal to `{\"__typename\":\"__EnumValue\",\"name\":\"EnumValue\",\"description\":\"Description\",\"isDeprecated\":true,\"deprecationReason\":\"DeprecationReason\"}`")
	}
}

func Test___EnumValue_f75cbd9d2b_with_error(t *testing.T) {
	ec := executionContext{
		OperationContext: &graphql.OperationContext{
			Variables: map[string]interface{}{},
		},
	}

	obj := &introspection.EnumValue{
		Name: "EnumValue",
		Description: "Description",
		IsDeprecated: true,
		DeprecationReason: "DeprecationReason",
	}

	sel := ast.SelectionSet{
		Selections: []ast.Selection{
			&ast.Field{
				Name: "__typename",
			},
			&ast.Field{
				Name: "name",
			},
			&ast.Field{
				Name: "description",
			},
			&ast.Field{
				Name: "isDeprecated",
			},
			&ast.Field{
				Name: "deprecationReason",
			},
			&ast.Field{
				Name: "unknown",
			},
		},
	}

	out := ec.___EnumValue(context.Background(), sel, obj)

	if out == nil {
		t.Error("out is nil")
		return
	}

	if out.MarshalGQL() != `{"__typename":"__EnumValue","name":"EnumValue","description":"Description","isDeprecated":true,"deprecationReason":"DeprecationReason"}` {
		t.Error("out.MarshalGQL() is not equal to `{\"__typename\":\"__EnumValue\",\"name\":\"EnumValue\",\"description\":\"Description\",\"isDeprecated\":true,\"deprecationReason\":\"DeprecationReason\"}`")
	}
}
