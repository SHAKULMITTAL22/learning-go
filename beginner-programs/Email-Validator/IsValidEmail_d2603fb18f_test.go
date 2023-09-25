package Validator

import (
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	testCases := []struct {
		email    string
		expected bool
	}{
		//Test case for valid email
		{"test@example.com", true},
		//Test case for invalid email
		{"test@", false},
		//Test case for email length > 254
		{"a@b.c" + string(make([]byte, 253)), false},
		//Test case for email without domain
		{"test@", false},
		//Test case for email without @ symbol
		{"testexample.com", false},
		//Test case for email with special characters in local part
		{"test..test@example.com", false},
		//Test case for email with special characters in domain part
		{"test@exa$mple.com", false},
	}

	for _, testCase := range testCases {
		if result := IsValidEmail(testCase.email); result != testCase.expected {
			t.Errorf("TestIsValidEmail failed with Email: %s. Got: %v, expected: %v", testCase.email, result, testCase.expected)
		}
	}
}
