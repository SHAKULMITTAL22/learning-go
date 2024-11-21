package generated

import (
	"context"
	"testing"
)

type testExecutionContext struct{}

func (ec *testExecutionContext) field_Query___type_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	return rawArgs, nil
}

func TestField_Query___type_args_5d7369acee(t *testing.T) {
	ec := &testExecutionContext{}

	t.Run("test with valid arguments", func(t *testing.T) {
		rawArgs := map[string]interface{}{
			"name": "test",
		}

		expected := map[string]interface{}{
			"name": "test",
		}

		result, err := ec.field_Query___type_args(context.Background(), rawArgs)
		if err != nil {
			t.Errorf("unexpected error: %v, args: %v", err, rawArgs)
		}

		if result["name"] != expected["name"] {
			t.Errorf("unexpected result: got %v, want %v, args: %v", result, expected, rawArgs)
		} else {
			t.Logf("success: expected result matches got result, args: %v", rawArgs)
		}
	})

	t.Run("test with invalid arguments", func(t *testing.T) {
		rawArgs := map[string]interface{}{
			"name": 123,
		}

		_, err := ec.field_Query___type_args(context.Background(), rawArgs)
		if err == nil {
			t.Errorf("expected error, but got nil, args: %v", rawArgs)
		} else {
			t.Logf("success: expected error and got error, args: %v", rawArgs)
		}
	})

	t.Run("test with missing arguments", func(t *testing.T) {
		rawArgs := map[string]interface{}{}

		expected := map[string]interface{}{
			"name": "",
		}

		result, err := ec.field_Query___type_args(context.Background(), rawArgs)
		if err != nil {
			t.Errorf("unexpected error: %v, args: %v", err, rawArgs)
		}

		if result["name"] != expected["name"] {
			t.Errorf("unexpected result: got %v, want %v, args: %v", result, expected, rawArgs)
		} else {
			t.Logf("success: expected result matches got result, args: %v", rawArgs)
		}
	})
}
