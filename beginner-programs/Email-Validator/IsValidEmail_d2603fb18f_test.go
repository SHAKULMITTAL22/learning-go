package Validator

import (
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
		t.Log("Test case 1 passed")
	}

	// Test case 2: Invalid email address
	email = "johndoe"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 2 passed")
	}

	// Test case 3: Email address with invalid characters
	email = "johndoe@example#$%^&*()"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 3 passed")
	}

	// Test case 4: Email address with a domain that doesn't exist
	email = "johndoe@example.xyz"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 4 passed")
	}

	// Test case 5: Email address with a local part that is too long
	email = "johndoe@example.com" + "a" * 64
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 5 passed")
	}

	// Test case 6: Email address with a domain part that is too long
	email = "johndoe@" + "example.com" * 64
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 6 passed")
	}

	// Test case 7: Email address with a missing "@" symbol
	email = "johndoexample.com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 7 passed")
	}

	// Test case 8: Email address with multiple "@" symbols
	email = "johndoe@example@example.com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 8 passed")
	}

	// Test case 9: Email address with a missing domain name
	email = "johndoe@"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 9 passed")
	}

	// Test case 10: Email address with a missing local part
	email = "@example.com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 10 passed")
	}

	// Test case 11: Email address with a space in the local part
	email = "john doe@example.com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 11 passed")
	}

	// Test case 12: Email address with a space in the domain part
	email = "johndoe@example. com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 12 passed")
	}

	// Test case 13: Email address with a comma in the local part
	email = "john,doe@example.com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 13 passed")
	}

	// Test case 14: Email address with a comma in the domain part
	email = "johndoe@example,com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 14 passed")
	}

	// Test case 15: Email address with a semicolon in the local part
	email = "john;doe@example.com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 15 passed")
	}

	// Test case 16: Email address with a semicolon in the domain part
	email = "johndoe@example;com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 16 passed")
	}

	// Test case 17: Email address with a backslash in the local part
	email = "john\\doe@example.com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 17 passed")
	}

	// Test case 18: Email address with a backslash in the domain part
	email = "johndoe@example\\com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 18 passed")
	}

	// Test case 19: Email address with a double quote in the local part
	email = "john\"doe@example.com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 19 passed")
	}

	// Test case 20: Email address with a double quote in the domain part
	email = "johndoe@example\"com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 20 passed")
	}

	// Test case 21: Email address with a single quote in the local part
	email = "john'doe@example.com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 21 passed")
	}

	// Test case 22: Email address with a single quote in the domain part
	email = "johndoe@example'com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 22 passed")
	}

	// Test case 23: Email address with a less than sign in the local part
	email = "john<doe@example.com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 23 passed")
	}

	// Test case 24: Email address with a less than sign in the domain part
	email = "johndoe@example<com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 24 passed")
	}

	// Test case 25: Email address with a greater than sign in the local part
	email = "john>doe@example.com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 25 passed")
	}

	// Test case 26: Email address with a greater than sign in the domain part
	email = "johndoe@example>com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 26 passed")
	}

	// Test case 27: Email address with a bracket in the local part
	email = "john[doe@example.com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 27 passed")
	}

	// Test case 28: Email address with a bracket in the domain part
	email = "johndoe@example]com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 28 passed")
	}

	// Test case 29: Email address with a curly brace in the local part
	email = "john{doe@example.com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 29 passed")
	}

	// Test case 30: Email address with a curly brace in the domain part
	email = "johndoe@example}com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 30 passed")
	}

	// Test case 31: Email address with a pipe in the local part
	email = "john|doe@example.com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 31 passed")
	}

	// Test case 32: Email address with a pipe in the domain part
	email = "johndoe@example|com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 32 passed")
	}

	// Test case 33: Email address with a tilde in the local part
	email = "john~doe@example.com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 33 passed")
	}

	// Test case 34: Email address with a tilde in the domain part
	email = "johndoe@example~com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 34 passed")
	}

	// Test case 35: Email address with a percent sign in the local part
	email = "john%doe@example.com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 35 passed")
	}

	// Test case 36: Email address with a percent sign in the domain part
	email = "johndoe@example%com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 36 passed")
	}

	// Test case 37: Email address with an ampersand in the local part
	email = "john&doe@example.com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 37 passed")
	}

	// Test case 38: Email address with an ampersand in the domain part
	email = "johndoe@example&com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 38 passed")
	}

	// Test case 39: Email address with a hash sign in the local part
	email = "john#doe@example.com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 39 passed")
	}

	// Test case 40: Email address with a hash sign in the domain part
	email = "johndoe@example#com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 40 passed")
	}

	// Test case 41: Email address with a question mark in the local part
	email = "john?doe@example.com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 41 passed")
	}

	// Test case 42: Email address with a question mark in the domain part
	email = "johndoe@example?com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 42 passed")
	}

	// Test case 43: Email address with an exclamation mark in the local part
	email = "john!doe@example.com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 43 passed")
	}

	// Test case 44: Email address with an exclamation mark in the domain part
	email = "johndoe@example!com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 44 passed")
	}

	// Test case 45: Email address with a dollar sign in the local part
	email = "john$doe@example.com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 45 passed")
	}

	// Test case 46: Email address with a dollar sign in the domain part
	email = "johndoe@example$com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 46 passed")
	}

	// Test case 47: Email address with a colon in the local part
	email = "john:doe@example.com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 47 passed")
	}

	// Test case 48: Email address with a colon in the domain part
	email = "johndoe@example:com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 48 passed")
	}

	// Test case 49: Email address with a semicolon in the local part
	email = "john;doe@example.com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 49 passed")
	}

	// Test case 50: Email address with a semicolon in the domain part
	email = "johndoe@example;com"
	expected = false
	actual = IsValidEmail(email)
	if actual != expected {
		t.Errorf("Expected %t, got %t for email %s", expected, actual, email)
	} else {
		t.Log("Test case 50 passed")
	}
}
