package Validator

import (
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	//Test case for valid email
	email := "test@example.com"
	expected := true
	if result := IsValidEmail(email); result != expected {
		t.Errorf("TestIsValidEmail failed with Email: %s. Got: %v, expected: %v", email, result, expected)
	} 

	//Test case for invalid email
	email = "test@"
	expected = false
	if result := IsValidEmail(email); result != expected {
		t.Errorf("TestIsValidEmail failed with Email: %s. Got: %v, expected: %v", email, result, expected)
	}

	//Test case for email length > 254
	email = "a@b.c" + string(make([]byte, 253))
	expected = false
	if result := IsValidEmail(email); result != expected {
		t.Errorf("TestIsValidEmail failed with Email: %s. Got: %v, expected: %v", email, result, expected)
	}
}
