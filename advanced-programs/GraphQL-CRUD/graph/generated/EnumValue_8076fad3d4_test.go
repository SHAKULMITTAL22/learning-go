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

func TestEnumValue_8076fad3d4(t *testing.T) {
	ec := executionContext{
		introspection.EnumValue{
			Name: "EnumValue",
			Description: "EnumValue description",
			DeprecationReason: "EnumValue deprecation reason",
		},
	}

	sel := ast.SelectionSet{
		Fields: []*ast.Field{
			{
				Name: "name",
			},
			{
				Name: "description",
			},
			{
				Name: "deprecationReason",
			},
		},
	}

	buf := bytes.NewBuffer(nil)
	err := ec.marshalN__EnumValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValue(context.Background(), sel, &ec.EnumValue).MarshalGQL(buf)
	if err != nil {
		t.Error(err)
	}

	expected := `{
  "name": "EnumValue",
  "description": "EnumValue description",
  "deprecationReason": "EnumValue deprecation reason"
}`

	if buf.String() != expected {
		t.Errorf("Expected %s, got %s", expected, buf.String())
	}
}

func TestEnumValue_error(t *testing.T) {
	ec := executionContext{
		introspection.EnumValue{
			Name: "EnumValue",
			Description: "EnumValue description",
			DeprecationReason: "EnumValue deprecation reason",
		},
	}

	sel := ast.SelectionSet{
		Fields: []*ast.Field{
			{
				Name: "name",
			},
			{
				Name: "description",
			},
			{
				Name: "deprecationReason",
			},
			{
				Name: "invalidField",
			},
		},
	}

	buf := bytes.NewBuffer(nil)
	err := ec.marshalN__EnumValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValue(context.Background(), sel, &ec.EnumValue).MarshalGQL(buf)
	if err == nil {
		t.Error("Expected error, got nil")
	}

	expected := `unknown field "invalidField"`

	if err.Error() != expected {
		t.Errorf("Expected %s, got %s", expected, err.Error())
	}
}
