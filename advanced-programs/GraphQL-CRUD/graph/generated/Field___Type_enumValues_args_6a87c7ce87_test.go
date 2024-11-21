package tests

import (
	"testing"
	"context"

	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/graph/model"
)

type executionContext struct{
	// TODO: Implement a proper execution context
}

func (ec *executionContext) unmarshalOBoolean2bool(ctx context.Context, v interface{}) (bool, error) {
	if tmp, ok := v.(bool); ok {
		return tmp, nil
	}
	return false, nil
}

func (ec *executionContext) field___Type_enumValues_args(ctx context.Context, field model.__Type) (bool, error) {
	var var0 bool

	if tmp, ok := field.EnumValues.(bool); ok {
		var0, _ = ec.unmarshalOBoolean2bool(ctx, tmp)
	} else {
		var0, _ = ec.unmarshalOBoolean2bool(ctx, tmp)
	}

	return var0, nil
}

// Testing method for the function 'field___Type_enumValues_args'
func TestField___Type_enumValues_args_6a87c7ce87(t *testing.T) {

	t.Run("Test Success Case", func(t *testing.T) {
		ctx := context.Background()

		field := model.__Type{ EnumValues: true }

		execCtx := &executionContext{}

		args, err := execCtx.field___Type_enumValues_args(ctx, field)

		if err != nil {
			t.Error("Expected no error, got", err)
			t.Log("Args Provided:", field.EnumValues)
		}

		if !args {
			t.Error("Expected includeDeprecated to be true, but got false")
		} else {
			t.Log("Test Success Case Passed.")
		}
	})

	t.Run("Test Failure Case", func(t *testing.T) {

		ctx := context.Background()

		field := model.__Type{ EnumValues: "foo" }

		execCtx := &executionContext{}

		args, err := execCtx.field___Type_enumValues_args(ctx, field)

		if err == nil {
			t.Error("Expected error, got none")
			t.Log("Args Provided:", field.EnumValues)
		}

		if args != false {
			t.Error("Expected no response, but got", args)
		} else {
			t.Log("Test Failure Case Passed.")
		}
	})

	// TODO: Write more test cases for other scenarios/edge cases/error handling
}

