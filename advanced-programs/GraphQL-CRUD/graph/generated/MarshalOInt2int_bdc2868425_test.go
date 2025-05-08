package generated

import (
	"context"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

func TestMarshalOInt2int_bdc2868425(t *testing.T) {
	ctx := context.Background()
	selectionSet := ast.SelectionSet{}

	ec := &executionContext{}

	t.Run("Positive Integer", func(t *testing.T) {
		v := 10

		marshalledValue := ec.marshalOInt2int(ctx, selectionSet, v)

		if marshalledValue == nil || marshalledValue.(graphql.Int) != graphql.Int(v) {
			t.Errorf("Got '%v' but expected '%v'", marshalledValue, v)
		} else {
			t.Logf("Got expected value '%v'", v)
		}
	})

	t.Run("Negative Integer", func(t *testing.T) {
		v := -10

		marshalledValue := ec.marshalOInt2int(ctx, selectionSet, v)

		if marshalledValue == nil || marshalledValue.(graphql.Int) != graphql.Int(v) {
			t.Errorf("Got '%v' but expected '%v'", marshalledValue, v)
		} else {
			t.Logf("Got expected value '%v'", v)
		}
	})

	t.Run("Zero Value", func(t *testing.T) {
		v := 0

		marshalledValue := ec.marshalOInt2int(ctx, selectionSet, v)

		if marshalledValue == nil || marshalledValue.(graphql.Int) != graphql.Int(v) {
			t.Errorf("Got '%v' but expected '%v'", marshalledValue, v)
		} else {
			t.Logf("Got expected value '%v'", v)
		}
	})
}
