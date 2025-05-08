package Validator

import (
	"regexp"
	"testing"
)

func TestIsValidEmail_d2603fb18f(t *testing.T) {
	// Test case 1: Valid email address
	email := "johndoe@example.com"
	expected := true
	actual := IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Logf("Test case 1 passed: Expected %t, got %t for email %s", expected, actual, email)
	}

	// Test case 2: Invalid email address
	email = "johndoe"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Logf("Test case 2 passed: Expected %t, got %t for email %s", expected, actual, email)
	}

	// Test case 3: Email address with invalid characters
	email = "johndoe@example#com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Logf("Test case 3 passed: Expected %t, got %t for email %s", expected, actual, email)
	}

	// Test case 4: Email address with spaces
	email = "johndoe@example. com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Logf("Test case 4 passed: Expected %t, got %t for email %s", expected, actual, email)
	}

	// Test case 5: Email address with a dot at the end
	email = "johndoe@example."
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Logf("Test case 5 passed: Expected %t, got %t for email %s", expected, actual, email)
	}

	// Test case 6: Email address with two dots in a row
	email = "johndoe@example..com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Logf("Test case 6 passed: Expected %t, got %t for email %s", expected, actual, email)
	}

	// Test case 7: Email address with a hyphen at the beginning
	email = "-johndoe@example.com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Logf("Test case 7 passed: Expected %t, got %t for email %s", expected, actual, email)
	}

	// Test case 8: Email address with a hyphen in the middle
	email = "john-doe@example.com"
	expected = true
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Logf("Test case 8 passed: Expected %t, got %t for email %s", expected, actual, email)
	}

	// Test case 9: Email address with a hyphen at the end
	email = "johndoe-@example.com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Logf("Test case 9 passed: Expected %t, got %t for email %s", expected, actual, email)
	}

	// Test case 10: Email address with an underscore at the beginning
	email = "_johndoe@example.com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Logf("Test case 10 passed: Expected %t, got %t for email %s", expected, actual, email)
	}

	// Test case 11: Email address with an underscore in the middle
	email = "john_doe@example.com"
	expected = true
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Logf("Test case 11 passed: Expected %t, got %t for email %s", expected, actual, email)
	}

	// Test case 12: Email address with an underscore at the end
	email = "johndoe_@example.com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Logf("Test case 12 passed: Expected %t, got %t for email %s", expected, actual, email)
	}

	// Test case 13: Email address with a plus sign at the beginning
	email = "+johndoe@example.com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Logf("Test case 13 passed: Expected %t, got %t for email %s", expected, actual, email)
	}

	// Test case 14: Email address with a plus sign in the middle
	email = "john+doe@example.com"
	expected = true
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Logf("Test case 14 passed: Expected %t, got %t for email %s", expected, actual, email)
	}

	// Test case 15: Email address with a plus sign at the end
	email = "johndoe+@example.com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Logf("Test case 15 passed: Expected %t, got %t for email %s", expected, actual, email)
	}
}
