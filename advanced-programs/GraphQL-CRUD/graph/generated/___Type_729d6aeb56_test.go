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

func Test___Type_729d6aeb56(t *testing.T) {
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
				Name: "kind",
			},
			{
				Name: "name",
			},
			{
				Name: "description",
			},
			{
				Name: "fields",
			},
			{
				Name: "interfaces",
			},
			{
				Name: "possibleTypes",
			},
			{
				Name: "enumValues",
			},
			{
				Name: "inputFields",
			},
			{
				Name: "ofType",
			},
		},
	}

	obj := &introspection.Type{
		Kind:        "SCALAR",
		Name:        "String",
		Description: "The `String` scalar type represents textual data, represented as UTF-8 character sequences. The String type is most often used by GraphQL to represent free-form human-readable text.",
	}

	out := ec.___Type(context.Background(), sel, obj)
	marshaler, ok := out.(graphql.Marshaler)
	if !ok {
		t.Error("expected out to be a graphql.Marshaler")
		return
	}

	buf := bytes.NewBuffer(nil)
	err := marshaler.MarshalGQL(buf)
	if err != nil {
		t.Error("expected MarshalGQL to succeed", err)
		return
	}

	expected := `{
  "__typename": "__Type",
  "kind": "SCALAR",
  "name": "String",
  "description": "The \\`String\\` scalar type represents textual data, represented as UTF-8 character sequences. The String type is most often used by GraphQL to represent free-form human-readable text.",
  "fields": null,
  "interfaces": null,
  "possibleTypes": null,
  "enumValues": null,
  "inputFields": null,
  "ofType": null
}`

	if buf.String() != expected {
		t.Errorf("expected output to be %q, got %q", expected, buf.String())
	}
}

func Test___Type_with_invalid_field(t *testing.T) {
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
				Name: "invalid_field",
			},
		},
	}

	obj := &introspection.Type{
		Kind:        "SCALAR",
		Name:        "String",
		Description: "The `String` scalar type represents textual data, represented as UTF-8 character sequences. The String type is most often used by GraphQL to represent free-form human-readable text.",
	}

	out := ec.___Type(context.Background(), sel, obj)
	marshaler, ok := out.(graphql.Marshaler)
	if !ok {
		t.Error("expected out to be a graphql.Marshaler")
		return
	}

	buf := bytes.NewBuffer(nil)
	err := marshaler.MarshalGQL(buf)
	if err != nil {
		t.Error("expected MarshalGQL to succeed", err)
		return
	}

	expected := `{
  "__typename": "__Type",
  "invalid_field": null
}`

	if buf.String() != expected {
		t.Errorf("expected output to be %q, got %q", expected, buf.String())
	}
}
