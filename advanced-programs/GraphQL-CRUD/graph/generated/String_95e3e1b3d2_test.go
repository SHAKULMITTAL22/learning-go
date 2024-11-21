package generated

import (
	"context"
	"testing"
)

func TestUnmarshalOString2ᚖstring(t *testing.T) {
	ec := &executionContext{}

	t.Run("Success Case: Valid String Argument", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("Panic: %v", r)
			}
		}()

		v := "test string"
		ctx := context.Background()
		res, err := ec.unmarshalOString2ᚖstring(ctx, v)
		if err != nil || *res != v {
			t.Error("Expected: ", v, ", but got ", *res)
		} else {
			t.Log("Success test passed!")
		}
	})

	t.Run("Case: Nil Argument", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("Panic: %v", r)
			}
		}()

		v := nil
		ctx := context.Background()
		res, err := ec.unmarshalOString2ᚖstring(ctx, v)
		if err != nil || res != nil {
			t.Error("Expected nil, but got ", res)
		} else {
			t.Log("Success test passed!")
		}
	})

	t.Run("Error Case: Invalid String Argument", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("Panic: %v", r)
			}
		}()

		v := 1234
		ctx := context.Background()
		_, err := ec.unmarshalOString2ᚖstring(ctx, v)
		if err == nil {
			t.Error("Expected error, but got nil")
		} else {
			t.Log("Error test passed!")
		}
	})
}
