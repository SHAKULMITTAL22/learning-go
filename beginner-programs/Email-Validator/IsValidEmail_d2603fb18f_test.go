package Validator

import (
	"testing"
)

func TestIsValidEmail_d2603fb18f(t *testing.T) {
	// Test case 1: Valid email address
	email := "example@example.com"
	expected := true
	actual := IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 1 passed")
	}

	// Test case 2: Invalid email address
	email = "example"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 2 passed")
	}

	// Test case 3: Email address with invalid characters
	email = "example@example.com!"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 3 passed")
	}

	// Test case 4: Email address with invalid domain
	email = "example@example"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 4 passed")
	}

	// Test case 5: Email address with invalid TLD
	email = "example@example.co"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 5 passed")
	}

	// Test case 6: Email address with valid TLD
	email = "example@example.com"
	expected = true
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 6 passed")
	}
}
