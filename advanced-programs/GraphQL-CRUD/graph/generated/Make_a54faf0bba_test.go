package generated

import (
    "context"
    "testing"

    "github.com/99designs/gqlgen/graphql"
    "github.com/99designs/gqlgen/graphql/introspection"
    "github.com/99designs/gqlgen/graphql/ast"
)

type executionContext struct{}

func (e *executionContext) marshalO__Field2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐFieldᚄ(ctx context.Context, sel ast.SelectionSet, fields []introspection.Field) graphql.Marshaler {
    // Marshal logic goes here for the function.
    return nil 
}

func TestMarshalOField2(t *testing.T) {
    // Creating execution context
    ec := &executionContext{}

    // Create a mock context
    ctx := context.Background()

    // Create a new instance of the introspection.field 
    field := &introspection.Field{
        Name: "MyField",
        Type: &introspection.TypeRef{
            Name: "MyType",
        },
    }

    // Create a selection set
    sel := ast.SelectionSet{}

    // The array of fields
    fields := []introspection.Field{*field}

    // Running the marshalO__Field2
    res := ec.marshalO__Field2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐFieldᚄ(ctx, sel, fields)

    // Check if the return type is non nil
    if res == nil {
        t.Error("Result is nil but it should not be")
    }
}

func TestMarshalOField2Nil(t *testing.T) {
    // Creating execution context
    ec := &executionContext{}

    // Create a mock context
    ctx := context.Background()

    // Create a new instance of the introspection.field
    field := introspection.Field{}

    // Create a selection set
    sel := ast.SelectionSet{}

    // The array of fields
    var fields []introspection.Field

    // Running the marshalO__Field2
    res := ec.marshalO__Field2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐFieldᚄ(ctx, sel, fields)

    // Check if the return type is non nil
    if res == graphql.Null {
        t.Error("For a nil field array,func should return graphql null")
    }
}

