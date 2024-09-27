package generated

import (
	"testing"
	"context"
)

// Test for unmarshalOInt2ᚖint function
func TestUnmarshalOInt2ᚖint(t *testing.T) {
	ec := &executionContext{} 

	t.Run("test for nil value", func(t *testing.T) {
		res, err := ec.unmarshalOInt2ᚖint(context.Background(), nil)
		if err != nil {
			t.Errorf("Received error when none was expected: %s", err)
		}
		if res != nil {
			t.Errorf("Result not as expected, expected nil, got: %v", res)
		}
		t.Log("Function completed successfully with nil input")
	})

	t.Run("test with valid int value", func(t *testing.T) {
		res, err := ec.unmarshalOInt2ᚖint(context.Background(), 123)
		if err != nil {
			t.Errorf("Received error when none was expected: %s", err)
		}
		if *res != 123 {
			t.Errorf("Result not as expected, expected 123, got: %v", *res)
		}
		t.Log("Function completed successfully with valid integer input")
	})

	t.Run("test with non-int value", func(t *testing.T) {
		_, err := ec.unmarshalOInt2ᚖint(context.Background(), "123")
		if err == nil {
			t.Error("An error was expected, but no error was returned")
		}
		t.Log("Function completed successfully with non-integer input, as an error was expected")
	})
}
